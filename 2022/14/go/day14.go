// Advent of code 2022, day 14
// https://adventofcode.com/2022/day/14
package main

import (
	"fmt"
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

// Define cavern 2d slice node as data type
type Node [2]int

// Cavern 2d vertical slice. Node map points to int specifying if node is
// blocked or not, and what is blocking
type Cavern struct {
	nodes       map[Node]int // -1 = spawn, 0 = air, 1 = rock, 2 = sand
	sandspot    Node
	settledSand int
	dim         [4]int // [xMin, xMax, zMin, zMax]
	height      int
	width       int
	floor       int
}

// Builds node map from scan input nodes. Draws lines between input nodes
// describing rock parts
func NewCavern(scan [][]Node, sandspot Node) *Cavern {
	// Keep track of total dimensions: [xMin, xMax, zMin, zMax]
	dim := [4]int{sandspot[0], sandspot[0], sandspot[1], sandspot[1]}
	nodes := map[Node]int{sandspot: -1}
	// Add rocks to cavern node map
	for _, rockLine := range scan {
		for i, rockNode := range rockLine {
			nodes[rockNode] = 1
			dim = updateDimensions(dim, rockNode)
			// Draw line if not first node in line
			if i == 0 {
				continue
			}
			for _, node := range generateLine(rockNode, rockLine[i-1]) {
				nodes[node] = 1
			}
		}
	}
	cavern := Cavern{
		nodes: nodes, sandspot: sandspot, dim: dim,
		height: dim[3] - dim[2], width: dim[1] - dim[0],
	}
	return &cavern
}

// Simulates dropping a node of sand.
// Returns boolean flagging if sand settled.
// If false, it will fall forever or is clogging spawn
func (cavern *Cavern) spawnSand() bool {
	loc := cavern.sandspot
	for {
		// Check for freefall
		down := Node{loc[0], loc[1] + 1}
		if cavern.floor == 0 && down[1] > cavern.dim[3] {
			return false
		}
		// Check down
		content := cavern.getContent(down)
		if content <= 0 {
			loc = down
			continue
		}
		// Check down-left
		downLeft := Node{loc[0] - 1, loc[1] + 1}
		content = cavern.getContent(downLeft)
		if content <= 0 {
			loc = downLeft
			continue
		}
		// Check down-right
		downRight := Node{loc[0] + 1, loc[1] + 1}
		content = cavern.getContent(downRight)
		if content <= 0 {
			loc = downRight
			continue
		}
		// Settle sand
		cavern.settledSand++
		cavern.nodes[loc] = 2
		cavern.dim = updateDimensions(cavern.dim, loc)
		// Check if sand settled at spawn
		return loc != cavern.sandspot
	}
}

// Returns content of target node
func (cavern *Cavern) getContent(node Node) int {
	// Check if floor
	if node[1] == cavern.floor {
		return 1
	}
	// Check map for content
	content, exists := cavern.nodes[node]
	if !exists {
		content = 0
	}
	return content
}

// Update cavern dimensions based on given node
func updateDimensions(dim [4]int, node Node) [4]int {
	x, z := node[0], node[1]
	if dim[0] > x {
		dim[0] = x
	} else if dim[1] < x {
		dim[1] = x
	}
	if dim[2] > z {
		dim[2] = z
	} else if dim[3] < z {
		dim[3] = z
	}
	return dim
}

// Returns a slice of nodes describing line between start and end
func generateLine(start Node, end Node) (nodes []Node) {
	x, z := start[0], start[1]
	dx, dz := x-end[0], z-end[1]
	for dx != 0 || dz != 0 {
		if dx > 0 {
			nodes = append(nodes, Node{x - 1, z})
			x--
			dx--
			continue
		} else if dx < 0 {
			nodes = append(nodes, Node{x + 1, z})
			x++
			dx++
			continue
		}
		if dz > 0 {
			nodes = append(nodes, Node{x, z - 1})
			z--
			dz--
			continue
		} else if dz < 0 {
			nodes = append(nodes, Node{x, z + 1})
			z++
			dz++
			continue
		}
	}
	return nodes
}

// Prints cavern map to console
func (cavern *Cavern) printMap() {
	charMap := map[int]string{
		-1: "+",
		0:  " ",
		1:  "#",
		2:  "o",
	}
	for z := cavern.dim[2]; z <= cavern.dim[3]; z++ {
		for x := cavern.dim[0]; x <= cavern.dim[1]; x++ {
			sign := " "
			if z == cavern.floor {
				sign = "#"
			} else {
				content, ok := cavern.nodes[Node{x, z}]
				if ok {
					sign = charMap[content]
				}
			}
			fmt.Print(sign)
		}
		fmt.Print("\n")
	}
	if cavern.floor >= 0 {
		for x := cavern.dim[0]; x <= cavern.dim[1]; x++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}

// Parses puzzle input from txt file.
// Returns 2D vertical slice of cavern
func readInput(path string) *Cavern {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")
	scan := [][]Node{}
	sandspot := Node{500, 0}
	// Generate rock node slices from input strings
	for _, line := range lines {
		nodes := []Node{}
		for _, nodeString := range strings.Split(line, " -> ") {
			split := strings.Split(nodeString, ",")
			x, _ := strconv.Atoi(split[0])
			z, _ := strconv.Atoi(split[1])
			nodes = append(nodes, Node{x, z})
		}
		scan = append(scan, nodes)
	}
	cavern := NewCavern(scan, sandspot)
	return cavern
}

// Part 1: How many units of sand come to rest before sand starts flowing into the abyss below?
func solvePart1(cavern *Cavern) int {
	// cavern.printMap()
	for cavern.spawnSand() {
	}
	// cavern.printMap()
	return cavern.settledSand
}

// Part 2: How many units of sand come to rest before sand is blocked?
func solvePart2(cavern *Cavern) int {
	// cavern.printMap()
	cavern.floor = cavern.dim[3] + 2
	for cavern.spawnSand() {
	}
	// cavern.printMap()
	return cavern.settledSand
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
	log.Printf("Units of sand at rest when freefall: %v\n", answer1)
	log.Printf("Units of sand at rest when cavern filled: %v\n", answer2)
}
