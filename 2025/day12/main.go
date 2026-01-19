package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 (example):", part1(loadInput("input.txt")))
}

func part1(shapes []Shape, regions []Region) int {
	result := 0

	for _, region := range regions {
		// Assume each shape is fully 3x3. If we have enough space, then it certainly fits
		numShapes := 0
		for _, count := range region.Shapes {
			numShapes += count
		}
		if (region.Length/3)*(region.Width/3) >= numShapes {
			result++
			continue
		}

		// If number of tiles in the shapes is greater than the max number of tiles in the region
		// then it certainly doesn't fit
		numTiles := 0
		for i, count := range region.Shapes {
			numTiles += shapes[i].NumTiles * count
		}
		if region.Width*region.Length < numTiles {
			continue
		}

		// Would need to check if the shapes fit the region with a shape-packing algorithm.
		// Turns out this is not necessary for the puzzle input :)
	}

	return result
}

type Shape struct {
	NumTiles int
	Grid     [][]string
}

type Region struct {
	Width  int
	Length int
	Shapes []int
}

func loadInput(filename string) ([]Shape, []Region) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	shapes := make([]Shape, 0)
	regions := make([]Region, 0)
	var currShape Shape

	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		if line == "" && len(regions) == 0 {
			// After empty line, add currShape to list of shapes
			shapes = append(shapes, currShape)
			currShape = Shape{}
		} else if strings.Contains(line, "#") || strings.Contains(line, ".") {
			// Building currShape
			currShape.Grid = append(currShape.Grid, strings.Split(line, ""))
			currShape.NumTiles += strings.Count(line, "#")
		} else if strings.Contains(line, "x") {
			// Parsing region
			parts := strings.Split(line, ":")
			dimensions := strings.Split(parts[0], "x")
			shapes := strings.TrimSpace(parts[1])

			w, _ := strconv.Atoi(dimensions[0])
			l, _ := strconv.Atoi(dimensions[1])

			var newRegion Region
			newRegion.Width = w
			newRegion.Length = l

			for _, count := range strings.Split(shapes, " ") {
				c, _ := strconv.Atoi(count)
				newRegion.Shapes = append(newRegion.Shapes, c)
			}

			regions = append(regions, newRegion)
		}
	}

	return shapes, regions
}
