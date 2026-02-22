package day10_2

import (
	"bufio"
	_ "embed"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Algorithm taken from https://old.reddit.com/r/adventofcode/comments/1pity70/2025_day_10_solutions/nuhhwl1/

type JoltageCounter []int

type Machine struct {
	buttons  [][]int
	joltages JoltageCounter
}

type ButtonCombination struct {
	joltages JoltageCounter
	presses  int
}

func NewJoltageCounter(size int) JoltageCounter {
	return make([]int, size)
}

func (counter JoltageCounter) isZero() bool {
	for i := 0; i < len(counter); i++ {
		if counter[i] != 0 {
			return false
		}
	}
	return true
}

func (counter JoltageCounter) smallerOrEqual(b JoltageCounter) bool {
	for i := 0; i < len(counter); i++ {
		if counter[i] > b[i] {
			return false
		}
	}
	return true
}

func (counter JoltageCounter) equalsModulo2(b JoltageCounter) bool {
	for i := 0; i < len(counter); i++ {
		if counter[i]%2 != b[i]%2 {
			return false
		}
	}
	return true
}

func (counter JoltageCounter) getMinButtonPresses(buttonCombinations []ButtonCombination) (int, bool) {
	if counter.isZero() {
		return 0, true
	}

	var res = math.MaxInt
	for _, comb := range buttonCombinations {
		if !comb.joltages.smallerOrEqual(counter) {
			continue
		}
		if !comb.joltages.equalsModulo2(counter) {
			continue
		}

		var nextCounter = NewJoltageCounter(len(counter))
		for i := 0; i < len(counter); i++ {
			nextCounter[i] = (counter[i] - comb.joltages[i]) / 2
		}
		rec, ok := nextCounter.getMinButtonPresses(buttonCombinations)
		if !ok {
			continue
		}

		if n := 2*rec + comb.presses; n < res {
			res = n
		}
	}
	if res < math.MaxInt {
		return res, true
	}
	return 0, false
}

func Run(inputFile string) int {
	var machines = parseMachines(inputFile)
	var totalButtonPresses int
	for _, m := range machines {
		var buttonCombinations = getAllButtonCombinations(m.buttons, len(m.joltages))
		var n, _ = m.joltages.getMinButtonPresses(buttonCombinations)
		totalButtonPresses += n
	}
	return totalButtonPresses
}

func parseMachines(inputFile string) []Machine {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var machines []Machine

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		var machine Machine

		for i := 0; i < len(parts); i++ {
			// ignore first part, not needed
			if i == 0 {
				continue
			}

			// last part is desired joltages
			if i == len(parts)-1 {
				joltages, err := parseIntList(strings.Trim(parts[i], "{}"))
				if err != nil {
					log.Fatal(err)
				}
				machine.joltages = joltages
				break
			}

			// middle parts are buttons
			button, err := parseIntList(strings.Trim(parts[i], "()"))
			if err != nil {
				log.Fatal(err)
			}
			machine.buttons = append(machine.buttons, button)
		}

		machines = append(machines, machine)
	}
	return machines
}

func parseIntList(intListString string) ([]int, error) {
	intStringList := strings.Split(intListString, ",")
	ints := make([]int, len(intStringList))
	for i, p := range intStringList {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		ints[i] = n
	}
	return ints, nil
}

func getAllButtonCombinations(buttons [][]int, m int) []ButtonCombination {
	var res []ButtonCombination

	var helper func(index int, currentCounter []int, presses int)
	helper = func(index int, currentCounter []int, presses int) {
		// Base case: we've made a decision for every button
		if index == len(buttons) {
			// Must copy the counter to avoid overwriting it in other branches
			finalCounter := make([]int, m)
			copy(finalCounter, currentCounter)
			res = append(res, ButtonCombination{finalCounter, presses})
			return
		}

		// OPTION 1: Don't press the current button
		helper(index+1, currentCounter, presses)

		// OPTION 2: Press the current button
		// Update the counter for the targets this button hits
		for _, targetIdx := range buttons[index] {
			currentCounter[targetIdx]++
		}
		helper(index+1, currentCounter, presses+1)

		// BACKTRACK: Undo the press so the next branch starts fresh
		for _, targetIdx := range buttons[index] {
			currentCounter[targetIdx]--
		}
	}

	// Start the process at the first button
	helper(0, make([]int, m), 0)
	return res
}
