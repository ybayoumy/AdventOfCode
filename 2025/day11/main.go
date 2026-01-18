package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Part 1 (example):", part1(loadInput("example.txt")))
	fmt.Println("Part 1 (input):", part1(loadInput("input.txt")))

	fmt.Println("Part 2 (example):", part2(loadInput("example2.txt")))
	fmt.Println("Part 2 (input):", part2(loadInput("input.txt")))
}

func part1(data map[string][]string) int {
	return findPaths(data, "you", "out", make(map[string]int))
}

func part2(data map[string][]string) int {
	path1 := findPaths(data, "svr", "dac", make(map[string]int)) *
		findPaths(data, "dac", "fft", make(map[string]int)) *
		findPaths(data, "fft", "out", make(map[string]int))

	path2 := findPaths(data, "svr", "fft", make(map[string]int)) *
		findPaths(data, "fft", "dac", make(map[string]int)) *
		findPaths(data, "dac", "out", make(map[string]int))

	return path1 + path2
}

func findPaths(data map[string][]string, from string, to string, cache map[string]int) int {
	if val, ok := cache[from]; ok {
		return val
	}

	if from == to {
		return 1
	}

	result := 0
	for _, next := range data[from] {
		result += findPaths(data, next, to, cache)
	}

	cache[from] = result

	return result
}

func loadInput(filename string) map[string][]string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	result := make(map[string][]string)

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

		parts := strings.Split(line, ":")
		outputs := strings.Split(strings.TrimSpace(parts[1]), " ")

		result[parts[0]] = outputs
	}

	return result
}
