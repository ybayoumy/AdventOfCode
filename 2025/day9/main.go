/*
Day 1: Secret Entrance

Source: https://adventofcode.com/2025/day/1
*/
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 (example):", part1(loadInput("example.txt")))
	fmt.Println("Part 1 (test data):", part1(loadInput("input.txt")))

	// fmt.Println("Part 2 (example):", part2("example.txt"))
	// fmt.Println("Part 2 (test data):", part2("input.txt"))
}

func part1(tiles []Tile) int {
	maxArea := 0

	for i := range tiles {
		for j := range tiles {
			if i == j {
				continue
			}

			a := calcArea(tiles[i], tiles[j])
			if a > maxArea {
				maxArea = a
			}
		}
	}

	return maxArea
}

type Tile struct {
	x, y int
}

func calcArea(t1 Tile, t2 Tile) int {
	xDiff := int(math.Abs(float64(t1.x-t2.x))) + 1
	yDiff := int(math.Abs(float64(t1.y-t2.y))) + 1

	return xDiff * yDiff
}

func loadInput(filename string) []Tile {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	tiles := make([]Tile, 0)

	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		tiles = append(tiles, Tile{x, y})
	}

	return tiles
}
