/*
Day 3: Lobby

Source: https://adventofcode.com/2025/day/3
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Part 1 (example):", part1(loadInput("example.txt")))
	fmt.Println("Part 1 (test data):", part1(loadInput("input.txt")))

	fmt.Println("Part 2 (example):", part2(loadInput("example.txt")))
	fmt.Println("Part 2 (test data):", part2(loadInput("input.txt")))
}

func part1(banks []string) int {
	var result int

	for _, bank := range banks {
		result += findLargestJoltage(bank, 2)
	}

	return result
}

func part2(banks []string) int {
	var result int

	for _, bank := range banks {
		result += findLargestJoltage(bank, 12)
	}

	return result
}

func loadInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return strings.Split(string(data), "\n")
}

func findLargestJoltage(bank string, digits int) int {
	var joltage int
	var nums []int

	bankLen := utf8.RuneCountInString(bank)
	for i := 0; i < bankLen; i++ {
		n, _ := strconv.Atoi(string(bank[i]))
		nums = append(nums, n)
	}

	lastHighestIndex := -1
	for i := digits - 1; i >= 0; i-- {
		var highest int
		var highestIndex int

		for j := lastHighestIndex + 1; j < bankLen-i; j++ {
			if nums[j] > highest {
				highest = nums[j]
				highestIndex = j
			}
		}

		joltage = joltage*10 + highest
		lastHighestIndex = highestIndex
	}

	return joltage
}
