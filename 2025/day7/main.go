/*
Day 7: Laboratories

Source: https://adventofcode.com/2025/day/7
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Part 1 (example):", part1(loadInput("example.txt")))
	fmt.Println("Part 1 (test data):", part1(loadInput("input.txt")))

	fmt.Println("Part 2 (example):", part2(loadInput("example.txt")))
	fmt.Println("Part 2 (test data):", part2(loadInput("input.txt")))
}

type Coordinate struct {
	i int
	j int
}

func part1(grid [][]string) int {
	numSplits := 0
	queue := []Coordinate{findStartCoord(&grid)}

	height := len(grid)
	width := len(grid[0])

	for len(queue) > 0 {
		currCoord := queue[0]
		queue = queue[1:]

		// reached the bottom of grid
		if currCoord.i == height-1 {
			continue
		}

		switch grid[currCoord.i+1][currCoord.j] {
		case ".":
			queue = append(queue, Coordinate{currCoord.i + 1, currCoord.j})
		case "^":
			numSplits++
			grid[currCoord.i+1][currCoord.j] = "-" // Replace the "^" with "-" so we don't count twice
			if currCoord.j < width-1 {
				queue = append(queue, Coordinate{currCoord.i + 1, currCoord.j + 1})
			}

			if currCoord.j > 0 {
				queue = append(queue, Coordinate{currCoord.i + 1, currCoord.j - 1})
			}
		}
	}

	return numSplits
}

func part2(grid [][]string) int {
	cache := make(map[Coordinate]int)
	return countPaths(findStartCoord(&grid), &grid, &cache)
}

func countPaths(coord Coordinate, grid *[][]string, cache *map[Coordinate]int) int {
	// using a cache to avoid recalculating branches
	if val, ok := (*cache)[coord]; ok {
		return val
	}

	height := len(*grid)
	width := len((*grid)[0])

	for i := coord.i; i < height-1; i++ {
		switch (*grid)[i][coord.j] {
		case ".":
			continue
		case "^":
			var rightCount, leftCount int
			if coord.j < width-1 {
				rightCount = countPaths(Coordinate{i + 1, coord.j + 1}, grid, cache)
			}

			if coord.j > 0 {
				leftCount = countPaths(Coordinate{i + 1, coord.j - 1}, grid, cache)
			}
			(*cache)[coord] = rightCount + leftCount
			return rightCount + leftCount
		}
	}

	//base case is only 1 path (reaches bottom of grid)
	return 1
}

func findStartCoord(grid *[][]string) Coordinate {
	for i := range *grid {
		for j := range (*grid)[i] {
			if (*grid)[i][j] == "S" {
				return Coordinate{i, j}
			}
		}
	}

	return Coordinate{-1, -1}
}

func loadInput(filename string) [][]string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	var grid [][]string
	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		split := strings.Split(line, "")
		grid = append(grid, split)
	}

	return grid
}
