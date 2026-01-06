/*
Day 2: Gift Shop

Source: https://adventofcode.com/2025/day/2
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
	fmt.Println("Part 1 (example):", part1("example.txt"))
	fmt.Println("Part 1 (test data):", part1("input.txt"))

	fmt.Println("Part 2 (example):", part2("example.txt"))
	fmt.Println("Part 2 (test data):", part2("input.txt"))
}

func part1(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var result int

	ranges := strings.Split(string(data), ",")
	for _, idRange := range ranges {
		ids := strings.Split(idRange, "-")

		firstId, _ := strconv.Atoi(ids[0])
		lastId, _ := strconv.Atoi(ids[1])

		for id := firstId; id <= lastId; id++ {
			if !isValidId(id) {
				result += id
			}
		}
	}

	return result
}

func part2(filename string) int {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var result int

	ranges := strings.Split(string(data), ",")
	for _, idRange := range ranges {
		ids := strings.Split(idRange, "-")

		firstId, _ := strconv.Atoi(ids[0])
		lastId, _ := strconv.Atoi(ids[1])

		for id := firstId; id <= lastId; id++ {
			if !isValidId2(id) {
				result += id
			}
		}
	}

	return result
}

// helpers for part1
func isValidId(id int) bool {
	idStr := strconv.Itoa(id)

	idLen := utf8.RuneCountInString(idStr)
	if idLen%2 != 0 {
		return true
	}

	if idStr[:idLen/2] == idStr[idLen/2:] {
		return false
	}
	return true
}

// helpers for part2
func isValidId2(id int) bool {
	idStr := strconv.Itoa(id)

	idLen := utf8.RuneCountInString(idStr)
	for i := 2; i <= idLen; i++ {
		parts := getEqualSplits(idStr, i)
		if len(parts) == 0 {
			continue
		}

		if allEquals(parts) {
			return false
		}
	}

	return true
}

func allEquals(s []string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			return false
		}
	}
	return true
}

func getEqualSplits(s string, n int) []string {
	var result []string

	sLen := utf8.RuneCountInString(s)
	if sLen%n != 0 {
		return result // empty slice
	}

	jump := sLen / n
	for i := 0; i < sLen; i += jump {
		result = append(result, s[i:i+jump])
	}

	return result
}
