/*
Day 1: Secret Entrance

Source: https://adventofcode.com/2025/day/1
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 (example):", part1("example.txt"))
	fmt.Println("Part 1 (test data):", part1("input.txt"))

	fmt.Println("Part 2 (example):", part2("example.txt"))
	fmt.Println("Part 2 (test data):", part2("input.txt"))
}

func part1(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read input file")
		os.Exit(1)
	}
	rotations := strings.Split(string(data), "\n")

	dial := 50
	result := 0
	for _, rt := range rotations {
		dir := rune(rt[0])
		num, _ := strconv.Atoi(string(rt[1:]))

		switch dir {
		case 'L':
			dial -= num % 100
			if dial < 0 {
				dial += 100
			}
		case 'R':
			dial += num % 100
			if dial > 99 {
				dial -= 100
			}
		}

		if dial == 0 {
			result += 1
		}
	}

	return result
}

func part2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to read input file")
		os.Exit(1)
	}
	rotations := strings.Split(string(data), "\n")

	dial := 50
	result := 0
	for _, rt := range rotations {
		dir := rune(rt[0])
		num, _ := strconv.Atoi(string(rt[1:]))

		prev_dial := dial

		switch dir {
		case 'L':
			dial -= num % 100
			if dial < 0 {
				dial += 100

				if prev_dial != 0 {
					result += 1
				}
			}
		case 'R':
			dial += num % 100
			if dial > 99 {
				dial -= 100
				if dial != 0 {
					result += 1
				}
			}
		}

		result += num / 100
		if dial == 0 {
			result += 1
		}
	}

	return result
}
