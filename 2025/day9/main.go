/*
Day 9: Movie Theater

Source: https://adventofcode.com/2025/day/9
*/
package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 (example):", part1(loadInput("example.txt")))
	fmt.Println("Part 1 (test data):", part1(loadInput("input.txt")))

	fmt.Println("Part 2 (example):", part2(loadInput("example.txt")))
	fmt.Println("Part 2 (test data):", part2(loadInput("input.txt")))
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

func part2(tiles []Tile) int {
	// Compress tile coordinates to represent easier in a grid
	yMap, xMap := compressCoords(tiles)

	// Generate the grid with tiles (#) filled in
	grid := generateGrid(tiles, yMap, xMap)

	// Find biggest recangle that doesn't interesect a "."
	maxArea := 0
	for i := range tiles {
		for j := range tiles {
			if i == j {
				continue
			}

			if isRectInside(grid, yMap[tiles[i].y], xMap[tiles[i].x], yMap[tiles[j].y], xMap[tiles[j].x]) {
				a := calcArea(tiles[i], tiles[j])
				if a > maxArea {
					maxArea = a
				}
			}
		}
	}

	return maxArea
}

func isRectInside(grid [][]string, y1, x1, y2, x2 int) bool {
	for x := min(x1, x2); x <= max(x1, x2); x++ {
		if grid[y1][x] == "." || grid[y2][x] == "." {
			return false
		}
	}

	for y := min(y1, y2); y <= max(y1, y2); y++ {
		if grid[y][x1] == "." || grid[y][x2] == "." {
			return false
		}
	}

	return true
}

func compressCoords(tiles []Tile) (map[int]int, map[int]int) {
	// sort by x
	sortedX := make([]Tile, len(tiles))
	copy(sortedX, tiles)
	slices.SortFunc(sortedX, func(a Tile, b Tile) int { return a.x - b.x })

	count := 1 // start at 1 to account for padding the grid
	xMap := make(map[int]int)
	for _, t := range sortedX {
		if _, exists := xMap[t.x]; !exists {
			xMap[t.x] = count
			count++
		}
	}

	// sort by y
	sortedY := make([]Tile, len(tiles))
	copy(sortedY, tiles)
	slices.SortFunc(sortedY, func(a Tile, b Tile) int { return a.y - b.y })

	count = 1 // reset to 1
	yMap := make(map[int]int)
	for _, t := range sortedY {
		if _, exists := yMap[t.y]; !exists {
			yMap[t.y] = count
			count++
		}
	}

	return yMap, xMap
}

func generateGrid(tiles []Tile, yMap map[int]int, xMap map[int]int) [][]string {
	// Make the grid (padding by 1 on all sides)
	grid := make([][]string, len(yMap)+2)
	for y := range grid {
		for range len(xMap) + 2 {
			grid[y] = append(grid[y], ".")
		}
	}

	// draw perimeter of tiles
	for i, tile := range tiles {
		grid[yMap[tile.y]][xMap[tile.x]] = "#"

		next := tiles[(i+1)%len(tiles)]

		// fill in line of # between tiles
		if next.y == tile.y {
			if next.x > tile.x {
				for x := xMap[tile.x] + 1; x < xMap[next.x]; x++ {
					grid[yMap[tile.y]][x] = "#"
				}
			} else {
				for x := xMap[tile.x] - 1; x > xMap[next.x]; x-- {
					grid[yMap[tile.y]][x] = "#"
				}
			}
		} else if next.x == tile.x {
			if next.y > tile.y {
				for y := yMap[tile.y] + 1; y < yMap[next.y]; y++ {
					grid[y][xMap[tile.x]] = "#"
				}
			} else {
				for y := yMap[tile.y] - 1; y > yMap[next.y]; y-- {
					grid[y][xMap[tile.x]] = "#"
				}
			}
		}
	}

	// find inside point and fill the inside with #s
	startY, startX := findInsidePoint(grid)
	floodFill(grid, startY, startX)

	return grid
}

func floodFill(grid [][]string, startY int, startX int) {
	height := len(grid)
	width := len(grid[0])

	dirs := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	queue := [][2]int{{startY, startX}}

	for len(queue) > 0 {
		cur := queue[0]
		y := cur[0]
		x := cur[1]
		queue = queue[1:]

		if y < 0 || y >= height || x < 0 || x >= width || grid[y][x] == "#" {
			continue
		}

		grid[y][x] = "#"
		for _, d := range dirs {
			queue = append(queue, [2]int{y + d[0], x + d[1]})
		}
	}
}

func findInsidePoint(grid [][]string) (int, int) {
	width := len(grid[0])

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != "." {
				continue
			}

			rightFound := false
			leftFound := false

			for i := x; i < width; i++ {
				if grid[y][i] == "#" {
					rightFound = true
					break
				}
			}

			for i := x; i > 0; i-- {
				if grid[y][i] == "#" {
					leftFound = true
					break
				}
			}

			if rightFound && leftFound {
				return y, x
			}
		}
	}

	return -1, -1
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
