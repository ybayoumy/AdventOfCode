/*
Day 10: Factory

Source: https://adventofcode.com/2025/day/10
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
	fmt.Println("Part 1 (example):", part1("example.txt"))
	fmt.Println("Part 1 (test data):", part1("input.txt"))
}

func part1(filename string) int {
	machines := loadInput(filename)

	result := 0
	for _, m := range machines {
		result += m.findFewestButtonPresses()
	}

	return result
}

type Machine struct {
	desired uint16
	buttons []uint16
	joltage []int
}

func (m Machine) findFewestButtonPresses() int {
	type QueueItem struct {
		state       uint16
		buttonsUsed []int
	}
	queue := make([]QueueItem, 0)

	for i, b := range m.buttons {
		if b == m.desired {
			return 1
		}

		queue = append(queue, QueueItem{b, []int{i}})
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		for i, b := range m.buttons {
			if slices.Contains(item.buttonsUsed, i) {
				continue
			}

			newState := b ^ item.state
			if newState == m.desired {
				return len(item.buttonsUsed) + 1
			}

			queue = append(queue, QueueItem{newState, append([]int{i}, item.buttonsUsed...)})
		}
	}

	return -1
}

func loadInput(filename string) []Machine {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(f)

	machines := make([]Machine, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err.Error())
			os.Exit(1)
		}

		machines = append(machines, parseMachineLine(line))
	}

	return machines
}

func parseMachineLine(line string) Machine {
	// Parse the desired indicator lights
	//Ex. [.##.]
	// '[' is always at the beginning of line
	indicators := strings.Split(strings.Split(line[1:], "]")[0], "")
	numIndicators := len(indicators)
	var desired uint16
	for i, c := range indicators {
		if c == "#" {
			desired += 1 << (numIndicators - 1 - i) // Shift the bit to the correct position
		}
	}

	// Parse the buttons
	// Ex. (0,1,2) (2,3) (0,4)
	buttons := strings.Split(line, "(")[1:]
	var buttonsBits []uint16
	for _, b := range buttons {
		b = strings.Split(b, ")")[0]

		var bits uint16
		nums := strings.Split(b, ",")
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			bits += 1 << (numIndicators - 1 - n)
		}

		buttonsBits = append(buttonsBits, bits)
	}

	// Parse joltage requirements
	// Ex. {3,5,4,7}
	var joltage []int
	j := strings.Split(line, "{")[1]
	j = strings.Split(j, "}")[0]
	nums := strings.Split(j, ",")
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		joltage = append(joltage, n)
	}

	return Machine{desired, buttonsBits, joltage}
}
