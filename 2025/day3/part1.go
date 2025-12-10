package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var result int

	banks := strings.Split(string(data), "\r\n")
	for _, bank := range banks {
		joltage, err := findLargestJoltage(bank, 2)
		if err != nil {
			fmt.Println(err.Error())
		}

		result += joltage
	}

	fmt.Println("Result:", result)
}

// Handles part1 and part2 (did this retroactively)
func findLargestJoltage(bank string, digits int) (int, error) {
	var joltage int
	var nums []int

	bankLen := utf8.RuneCountInString(bank)
	for i := 0; i < bankLen; i++ {
		n, err := strconv.Atoi(string(bank[i]))
		if err != nil {
			return 0, nil
		}
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

	return joltage, nil
}
