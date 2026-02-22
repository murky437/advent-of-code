package day10

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	desiredLights []rune
	buttons       [][]int
}

func (m *Machine) GetEmptyLights() []rune {
	return []rune(strings.Repeat(".", len(m.desiredLights)))
}

func Run(inputFile string) int64 {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var machines []Machine

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var machine Machine

		for i := 0; i < len(parts); i++ {
			if i == 0 {
				machine.desiredLights = []rune(strings.Trim(parts[i], "[]"))
				continue
			}
			if i == len(parts)-1 {
				// ignore last part, not needed
				break
			}
			button, err := parseButton(parts[i])
			if err != nil {
				log.Fatal(err)
			}
			machine.buttons = append(machine.buttons, button)
		}

		machines = append(machines, machine)
	}

	var totalButtonPresses int64

	for i := range machines {
		totalButtonPresses += minButtonPresses(&machines[i])
	}

	return totalButtonPresses
}

func parseButton(buttonString string) ([]int, error) {
	buttonString = strings.Trim(buttonString, "()")
	buttonParts := strings.Split(buttonString, ",")
	button := make([]int, len(buttonParts))
	for j, p := range buttonParts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		button[j] = n
	}
	return button, nil
}

type ButtonPressVariation struct {
	state   []rune
	presses int64
}

func minButtonPresses(machine *Machine) int64 {
	queue := []ButtonPressVariation{{
		state:   append([]rune{}, machine.GetEmptyLights()...),
		presses: 0,
	}}

	visited := map[string]bool{string(machine.GetEmptyLights()): true}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if string(curr.state) == string(machine.desiredLights) {
			return curr.presses
		}

		for _, button := range machine.buttons {
			newState := append([]rune{}, curr.state...)
			for _, idx := range button {
				if newState[idx] == '.' {
					newState[idx] = '#'
				} else {
					newState[idx] = '.'
				}
			}
			key := string(newState)
			if !visited[key] {
				visited[key] = true
				queue = append(queue, ButtonPressVariation{newState, curr.presses + 1})
			}
		}
	}

	return math.MaxInt64
}
