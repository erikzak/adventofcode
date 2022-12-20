// Advent of code 2022, day 18
// https://adventofcode.com/2022/day/18
package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Node type for 3D grid coordinates
type Node [3]int

// Keeps track of droplet nodes, with methods for solving puzzles
type Droplet struct {
	nodes map[Node]int // 0 = undefined, 1 = rock, 2 = internal air, 3 = external air
	max   [3]int
}

func NewDroplet(nodes map[Node]int) *Droplet {
	max := Node{0, 0, 0}
	for node := range nodes {
		if node[0] > max[0] {
			max[0] = node[0]
		}
		if node[1] > max[1] {
			max[1] = node[1]
		}
		if node[2] > max[2] {
			max[2] = node[2]
		}
	}
	droplet := Droplet{nodes: nodes, max: max}
	return &droplet
}

// Returns content at node as int: 0 = undefined, 1 = rock, 2 = internal air, 3 = external air
func (droplet *Droplet) getContent(node Node) int {
	cube, found := droplet.nodes[node]
	if found {
		return cube
	}
	return 0
}

// Returns slice of neighbor nodes
func (droplet *Droplet) getNeighbors(node Node) []Node {
	deltas := []Node{
		{-1, 0, 0}, {1, 0, 0},
		{0, -1, 0}, {0, 1, 0},
		{0, 0, -1}, {0, 0, 1},
	}
	neighbors := []Node{}
	for _, delta := range deltas {
		node := Node{node[0] + delta[0], node[1] + delta[1], node[2] + delta[2]}
		neighbors = append(neighbors, node)
	}
	return neighbors
}

// Checks if air pocket node is interior or exterior
func (droplet *Droplet) isInteriorAir(node Node) bool {
	// Uses breadth-first search to check if outside is reachable
	queue := []Node{node}
	explored := map[Node]struct{}{node: {}}
	for {
		node := queue[0]
		if node[0] < 0 || node[0] > droplet.max[0] ||
			node[1] < 0 || node[1] > droplet.max[1] ||
			node[2] < 0 || node[2] > droplet.max[2] {
			return false
		}
		queue = queue[1:]
		for _, neighbor := range droplet.getNeighbors(node) {
			_, checked := explored[neighbor]
			if checked || droplet.getContent(neighbor) == 1 {
				continue
			}
			explored[neighbor] = struct{}{}
			queue = append(queue, neighbor)
		}
		if len(queue) == 0 {
			break
		}
	}
	// Is internal air pocket. Set all explored node content to 2
	for node := range explored {
		droplet.nodes[node] = 2
	}
	return true
}

// Inspects neighbors of given node. Returns number of unconnected sides
func (droplet *Droplet) checkSides(node Node, includeInterior bool) int {
	unconnectedSides := 0
	for _, neighbor := range droplet.getNeighbors(node) {
		content := droplet.getContent(neighbor)
		// Part 1, just check for unconnected sides
		if includeInterior {
			if content != 1 {
				unconnectedSides++
			}
			continue
		}

		// Part 2, ignore internal air pockets
		if content == 0 {
			isInterior := droplet.isInteriorAir(neighbor)
			if isInterior {
				content = 2
			} else {
				content = 3
			}
		}
		if content == 3 {
			unconnectedSides++
		}
	}
	return unconnectedSides
}

// Parses puzzle input from txt file.
// Returns droplet instace
func readInput(path string) *Droplet {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")
	nodes := map[Node]int{}
	for _, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		node := Node{x, y, z}
		nodes[node] = 1
	}
	droplet := NewDroplet(nodes)
	return droplet
}

// Part 1: What is the surface area of your scanned lava droplet?
func solvePart1(droplet *Droplet) int {
	unconnectedSides := 0
	for node := range droplet.nodes {
		unconnectedSides += droplet.checkSides(node, true)
	}
	return unconnectedSides
}

// Part 2: What is the exterior surface area of your scanned lava droplet?
func solvePart2(droplet *Droplet) int {
	unconnectedSides := 0
	for node := range droplet.nodes {
		unconnectedSides += droplet.checkSides(node, false)
	}
	return unconnectedSides
}

// Solves puzzle parts. Split up for benchmarking
func solvePuzzle() (int, int) {
	input := readInput(inputPath)
	answer1 := solvePart1(input)
	input = readInput(inputPath)
	answer2 := solvePart2(input)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Surface area of droplet: %v\n", answer1)
	log.Printf("Exterior surface area of droplet: %v\n", answer2)
}
