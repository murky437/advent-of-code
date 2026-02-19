package day5_2

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run(inputFile string) int64 {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ranges []Range

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		startAndStop := strings.Split(line, "-")
		start, err := strconv.ParseInt(startAndStop[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		stop, err := strconv.ParseInt(startAndStop[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, Range{start, stop})
	}

	var numOfFreshIngredientsIds int64

	mergedRanges := MergeRanges(ranges)

	for _, r := range mergedRanges {
		numOfFreshIngredientsIds += r.End - r.Start + 1
	}

	return numOfFreshIngredientsIds
}

type Range struct {
	Start int64
	End   int64
}

func MergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return nil
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{ranges[0]}

	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]

		if r.Start <= last.End+1 {
			if r.End > last.End {
				last.End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}
