/*
Day 6: Trash Compactor

Source: https://adventofcode.com/2025/day/6
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 (example):", part1(loadInput("example.txt")))
	fmt.Println("Part 1 (test data):", part1(loadInput("input.txt")))

	fmt.Println("Part 2 (example):", part2(loadInput("example.txt")))
	fmt.Println("Part 2 (test data):", part2(loadInput("input.txt")))
}

func part1(lines []string) int {
	data := make([][]string, 0, len(lines))
	for _, line := range lines {
		data = append(data, strings.Fields(line))
	}

	result := 0

	for j := 0; j < len(data[0]); j++ {
		nums := make([]int, 0, len(data))
		for i := 0; i < len(data); i++ {
			switch data[i][j] {
			case "*":
				result += multiplyNums(nums)
			case "+":
				result += addNums(nums)
			default:
				x, _ := strconv.Atoi(data[i][j])
				nums = append(nums, x)
			}
		}
	}

	return result
}

func part2(lines []string) int {
	result := 0

	nums := []int{}
	var operation rune
	for j := range lines[0] {
		numBuilder := []byte{}
		for i := range lines {
			switch lines[i][j] {
			case ' ':
				continue
			case '*':
				operation = '*'
			case '+':
				operation = '+'
			default:
				numBuilder = append(numBuilder, lines[i][j])
			}
		}

		if len(numBuilder) > 0 {
			x, _ := strconv.Atoi(string(numBuilder))
			nums = append(nums, x)
		} else {
			// full column of whitespace. calculate the value and reset nums
			switch operation {
			case '*':
				result += multiplyNums(nums)
			case '+':
				result += addNums(nums)
			}
			nums = []int{}
		}
	}

	if len(nums) > 0 {
		switch operation {
		case '*':
			result += multiplyNums(nums)
		case '+':
			result += addNums(nums)
		}
	}

	return result
}

func addNums(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func multiplyNums(nums []int) int {
	result := 1
	for _, num := range nums {
		result *= num
	}
	return result
}

func loadInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	lines := make([]string, 0)

	r := bufio.NewReader(f)
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
		lines = append(lines, line)
	}

	return lines
}
