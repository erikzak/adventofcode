// Advent of code 2022, day 12
// https://adventofcode.com/2022/day/12
package main

import (
	"log"
	"os"
	"strings"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Declares node type as basic 2-length array
type Node [2]int

// Keeps track of terrain grid, start and end nodes
type Terrain struct {
	width  int
	height int
	rows   []string
	start  Node
	end    Node
}

func NewTerrain(rows []string, start Node, end Node) Terrain {
	terrain := Terrain{rows: rows, start: start, end: end}
	terrain.width = len(rows[0])
	terrain.height = len(rows)
	return terrain
}

// A* finds a path from start to goal. Returns least number of steps.
// Adapted from https://en.wikipedia.org/wiki/A*_search_algorithm
func (terrain Terrain) AStar() (leastSteps int) {
	openSet := map[Node]struct{}{
		terrain.start: {},
	}
	cameFrom := map[Node]Node{}
	gScore := map[Node]int{}
	gScore[terrain.start] = 0
	fScore := map[Node]int{}
	fScore[terrain.start] = terrain.estimateDistanceToEnd(terrain.start)

	for {
		current := Node{-1, -1}
		lowestF := -1
		for node := range openSet {
			if lowestF == -1 || lowestF > fScore[node] {
				current, lowestF = node, fScore[node]
			}
		}

		if current == terrain.end {
			leastSteps = len(reconstructPath(cameFrom, current)) - 1
			return leastSteps
		}

		delete(openSet, current)
		neighbors := terrain.getNeighbors(current)
		for _, neighbor := range neighbors {
			tentativeGScore, cOk := gScore[current]
			if !cOk {
				tentativeGScore = 9999999
			} else {
				tentativeGScore++
			}
			neighborGScore, nOk := gScore[neighbor]
			if !nOk {
				neighborGScore = 9999999
			}
			if tentativeGScore < neighborGScore {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + terrain.estimateDistanceToEnd(neighbor)
				openSet[neighbor] = struct{}{}
			}
		}
		if len(openSet) == 0 {
			break
		}
	}

	return -1
}

// Returns slice of valid neighbors for the given terrain node
func (terrain Terrain) getNeighbors(node Node) (neighbors []Node) {
	neighbors = []Node{}
	nodeHeight := terrain.rows[node[1]][node[0]]
	neighborIndexes := []Node{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for _, potential := range neighborIndexes {
		target := Node{
			node[0] + potential[0],
			node[1] + potential[1],
		}
		if target[0] < 0 || target[0] >= terrain.width ||
			target[1] < 0 || target[1] >= terrain.height {
			continue
		}
		targetHeight := terrain.rows[target[1]][target[0]]
		if nodeHeight >= targetHeight-1 {
			neighbors = append(neighbors, target)
		}
	}
	return neighbors
}

// Returns manhattan distance between given node and end
func (terrain Terrain) estimateDistanceToEnd(node Node) int {
	return Abs(terrain.end[0]-node[0]) + Abs(terrain.end[1]-node[1])
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Reconstructs path from start to current node
func reconstructPath(cameFrom map[Node]Node, current Node) []Node {
	totalPath := []Node{current}
	for {
		next, ok := cameFrom[current]
		if ok {
			delete(cameFrom, current)
			totalPath = append([]Node{next}, totalPath...)
			current = totalPath[0]
			continue
		}
		break
	}
	return totalPath
}

// Parses puzzle input from txt file.
// Returns terrain instance
func readInput(path string) (terrain Terrain) {
	inputBytes, err := os.ReadFile(path)
	check(err)
	rows := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")
	start, end := Node{}, Node{}
	// Grab start and end indexes from terrain model and replace with heights
	for y, row := range rows {
		for x, char := range row {
			if char == 'S' {
				start = Node{x, y}
				rows[y] = strings.Replace(rows[y], "S", "a", 1)
			} else if char == 'E' {
				end = Node{x, y}
				rows[y] = strings.Replace(rows[y], "E", "z", 1)
			}
		}
	}
	terrain = NewTerrain(rows, start, end)
	return terrain
}

// Part 1: What is the fewest steps required to move from your current
// position to the location that should get the best signal?
func solvePart1(terrain Terrain) int {
	return terrain.AStar()
}

// Part 2: What is the fewest steps required to move starting from any square
// with elevation a to the location that should get the best signal?
func solvePart2(terrain Terrain) int {
	leastSteps := -1
	for y, row := range terrain.rows {
		for x, char := range row {
			if char == 'a' {
				terrain.start = Node{x, y}
				steps := terrain.AStar()
				if steps != -1 && (leastSteps == -1 || steps < leastSteps) {
					leastSteps = steps
				}
			}
		}
	}
	return leastSteps
}

// Solves puzzle parts. Split up for benchmarking
func solvePuzzle() (int, int) {
	input := readInput(inputPath)
	answer1 := solvePart1(input)
	answer2 := solvePart2(input)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Fewest steps to best signal: %v\n", answer1)
	log.Printf("Fewest steps from any a to best signal: %v\n", answer2)
}
