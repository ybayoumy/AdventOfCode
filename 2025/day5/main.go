/*
Day 5: Cafeteria

Source: https://adventofcode.com/2025/day/5
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	exampleRanges, exampleIds := loadInput("example.txt")
	testRanges, testIds := loadInput("input.txt")

	fmt.Println("Part 1 (example):", part1(exampleRanges, exampleIds))
	fmt.Println("Part 1 (test data):", part1(testRanges, testIds))

	fmt.Println("Part 2 (example):", part2(exampleRanges))
	fmt.Println("Part 2 (test data):", part2(testRanges))
}

func part1(ranges [][2]int, ids []int) int {
	numFresh := 0
	for _, id := range ids {
		if isFresh(id, ranges) {
			numFresh += 1
		}
	}

	return numFresh
}

func part2(ranges [][2]int) int {
	merged := mergeRanges(ranges)

	numFresh := 0
	for _, r := range merged {
		numFresh += r[1] - r[0] + 1
	}

	return numFresh
}

// part 1 helper
func isFresh(id int, ranges [][2]int) bool {
	for _, r := range ranges {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}
	return false
}

// part 2 helper
func mergeRanges(ranges [][2]int) [][2]int {
	slices.SortFunc(ranges, func(a [2]int, b [2]int) int {
		return a[0] - b[0]
	})

	var merged [][2]int
	for {
		if len(ranges) == 1 {
			merged = append(merged, ranges...)
			break
		}

		r1 := ranges[0]
		r2 := ranges[1]
		ranges = ranges[2:]

		if r2[0] > r1[1] {
			// no intersection
			merged = append(merged, r1)              // r1 range is finalized
			ranges = append([][2]int{r2}, ranges...) // r2 added prepended back to ranges
		} else {
			// they intersect
			// add combined range to the front of ranges
			ranges = append([][2]int{{min(r1[0], r2[0]), max(r1[1], r2[1])}}, ranges...)
		}
	}

	return merged
}

func loadInput(filename string) ([][2]int, []int) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	ranges := make([][2]int, 0)

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}

		line = strings.Trim(line, "\n")
		if line == "" {
			break
		}

		split := strings.Split(line, "-")
		lower, _ := strconv.Atoi(split[0])
		upper, _ := strconv.Atoi(split[1])

		ranges = append(ranges, [2]int{lower, upper})
	}

	ids := make([]int, 0)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}
		line = strings.Trim(line, "\n")

		num, _ := strconv.Atoi(line)
		ids = append(ids, num)
	}

	return ranges, ids
}
