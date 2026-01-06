/*
Day 4: Printing Department

Source: https://adventofcode.com/2025/day/4
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

func part1(grid [][]string) int {
	var result int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != "@" {
				continue
			}

			if numNeighbors(&grid, i, j) < 4 {
				result += 1
			}
		}
	}

	return result
}

func part2(grid [][]string) int {
	numRemoved := removeRolls(&grid)
	result := numRemoved

	for numRemoved > 0 {
		numRemoved = removeRolls(&grid)
		result += numRemoved
	}

	return result
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

func numNeighbors(grid *[][]string, row, col int) int {
	dirs := [][2]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, 1},  // Up + right
		{1, 1},   // Down + right
		{-1, -1}, // Up + left
		{1, -1},  // Down + left
	}
	var numNeighbors int

	for _, dir := range dirs {
		x := row + dir[0]
		y := col + dir[1]
		if x >= 0 && x < len(*grid) && y >= 0 && y < len((*grid)[x]) && (*grid)[x][y] == "@" {
			numNeighbors += 1
		}
	}

	return numNeighbors
}

func removeRolls(grid *[][]string) int {
	var numRemoved int

	for i := range *grid {
		for j := range (*grid)[i] {
			if (*grid)[i][j] != "@" {
				continue
			}

			if numNeighbors(grid, i, j) < 4 {
				numRemoved += 1
				(*grid)[i][j] = "."
			}
		}
	}

	return numRemoved
}
