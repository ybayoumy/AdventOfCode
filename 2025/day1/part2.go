package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Unable to read input file")
	}

	dial := 50
	result := 0
	lines := strings.Split(string(data), "\r\n")
	for _, line := range lines {
		dir := rune(line[0])
		num, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			fmt.Println(err.Error())
		}

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

	fmt.Println("Result:", result)
}
