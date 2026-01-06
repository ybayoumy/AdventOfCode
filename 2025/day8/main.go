/*
Day 8: Playground

Source: https://adventofcode.com/2025/day/8
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
	fmt.Println("Part 1 (example):", part1(loadInput("example.txt"), 10))
	fmt.Println("Part 1 (test data):", part1(loadInput("input.txt"), 1000))

	fmt.Println("Part 2 (example):", part2(loadInput("example.txt")))
	fmt.Println("Part 2 (test data):", part2(loadInput("input.txt")))
}

func part1(boxes []*JunctionBox, numPairs int) int {
	// calculate distances beween boxes and sort for lowest distance
	distances := computeDistances(boxes)

	// Each box is its own circuit to start
	circuits := initCircuits(boxes)

	// Find closest n pairings and combine their circuits
	for i := 0; i < numPairs; i++ {
		if i >= len(distances) {
			break
		}

		d := distances[i]
		d.box1.circuit.merge(d.box2.circuit)
	}

	// sort circuits in descending order by number of boxes in it
	slices.SortFunc(circuits, func(a *Circuit, b *Circuit) int {
		return len(b.boxes) - len(a.boxes)
	})

	// multiply top 3 and return result
	return len(circuits[0].boxes) * len(circuits[1].boxes) * len(circuits[2].boxes)
}

func part2(boxes []*JunctionBox) int {
	// calculate distances beween boxes and sort for lowest distance
	distances := computeDistances(boxes)

	// Each box is its own circuit to start
	initCircuits(boxes)

	// Keep combining circuits until all boxes are in a single circuit (or we run out of distances)
	for i := 0; i < len(distances); i++ {
		d := distances[i]
		d.box1.circuit.merge(d.box2.circuit)

		if len(d.box1.circuit.boxes) == len(boxes) {
			return d.box1.x * d.box2.x
		}
	}

	return -1
}

type JunctionBox struct {
	x, y, z int
	circuit *Circuit
}

func (box *JunctionBox) getDistance(other *JunctionBox) float64 {
	xDiff := box.x - other.x
	yDiff := box.y - other.y
	zDiff := box.z - other.z
	return math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff))
}

type Circuit struct {
	boxes map[*JunctionBox]struct{} // set
}

func (ct *Circuit) addBox(box *JunctionBox) {
	ct.boxes[box] = struct{}{}
	box.circuit = ct
}

func (ct *Circuit) merge(other *Circuit) {
	if other == ct {
		return
	}

	for box := range other.boxes {
		ct.addBox(box)
	}

	other.boxes = map[*JunctionBox]struct{}{} // empty the other circuit
}

type Distance struct {
	box1, box2 *JunctionBox
	distance   float64
}

func computeDistances(boxes []*JunctionBox) []Distance {
	distances := []Distance{}
	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			distances = append(distances, Distance{boxes[i], boxes[j], boxes[i].getDistance(boxes[j])})
		}
	}
	slices.SortFunc(distances, func(a Distance, b Distance) int {
		if a.distance-b.distance < 0 {
			return -1
		}
		return 1
	})

	return distances
}

func initCircuits(boxes []*JunctionBox) []*Circuit {
	circuits := make([]*Circuit, 0, len(boxes))
	for _, box := range boxes {
		ct := &Circuit{make(map[*JunctionBox]struct{})}
		ct.addBox(box)
		circuits = append(circuits, ct)
	}

	return circuits
}

func loadInput(filename string) []*JunctionBox {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var result []*JunctionBox
	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		result = append(result, &JunctionBox{x, y, z, nil})
	}

	return result
}
