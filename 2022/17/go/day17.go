// Advent of code 2022, day 17
// https://adventofcode.com/2022/day/17
package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Node type for 2D grid coordinates
type Node [2]int

// Define collection of shapes, in order, along with jet pattern in cave
type Tetris struct {
	shapes        []*Shape
	shpIdx        int
	jets          []rune
	jetIdx        int
	width         int
	chamber       map[Node]struct{} // Keeps track of filled nodes in chamber
	blocks        int
	currentHeight int // Height of top node
}

func NewTetris(jets []rune, width int) *Tetris {
	tetris := Tetris{jets: jets, width: width, chamber: map[Node]struct{}{}}
	minus := []Node{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	tetris.shapes = append(tetris.shapes, NewShape(minus, 4, 1))
	plus := []Node{{1, 2}, {0, 1}, {1, 1}, {2, 1}, {1, 0}}
	tetris.shapes = append(tetris.shapes, NewShape(plus, 3, 3))
	arrow := []Node{{2, 2}, {2, 1}, {0, 0}, {1, 0}, {2, 0}}
	tetris.shapes = append(tetris.shapes, NewShape(arrow, 3, 3))
	straight := []Node{{0, 3}, {0, 2}, {0, 1}, {0, 0}}
	tetris.shapes = append(tetris.shapes, NewShape(straight, 1, 4))
	square := []Node{{0, 1}, {1, 1}, {0, 0}, {1, 0}}
	tetris.shapes = append(tetris.shapes, NewShape(square, 2, 2))
	return &tetris
}

// Returns the next shape in slice of shapes
func (tetris *Tetris) getShape() *Shape {
	shape := tetris.shapes[tetris.shpIdx]
	tetris.shpIdx++
	if tetris.shpIdx == len(tetris.shapes) {
		tetris.shpIdx = 0
	}
	shape.position = [2]int{2, tetris.currentHeight + 4}
	return shape
}

// Returns the next jet in slice of jet patterns
func (tetris *Tetris) getJet() rune {
	jet := tetris.jets[tetris.jetIdx]
	tetris.jetIdx++
	if tetris.jetIdx == len(tetris.jets) {
		tetris.jetIdx = 0
	}
	return jet
}

// Returns content of chamber node. True if filled, false if empty
func (tetris *Tetris) isFilled(node Node) bool {
	if node[0] < 0 || node[0] >= tetris.width {
		return true
	}
	_, filled := tetris.chamber[node]
	return filled
}

// Moves shape left or right, if possible
func (tetris *Tetris) move(shape *Shape, direction rune) {
	dx := 1
	if direction == '<' {
		dx = -1
	}
	shape.position[0] += dx
	for node := range shape.getChamberNodes() {
		if tetris.isFilled(node) {
			shape.position[0] -= dx
			break
		}
	}
}

// Drops shape one row if possible, returns bool flagging if shape is at rest
func (tetris *Tetris) drop(shape *Shape) bool {
	shape.position[1]--
	for node := range shape.getChamberNodes() {
		if tetris.isFilled(node) {
			shape.position[1]++
			return true
		}
		if node[1] == 0 {
			shape.position[1]++
			return true
		}
	}
	return false
}

// Draws chamber
func (tetris *Tetris) draw(shape *Shape) {
	shapeNodes := shape.getChamberNodes()
	height := shape.position[1] + shape.height - 1
	if height == -1 {
		height = tetris.currentHeight
	}
	for y := height; y > 0; y-- {
		fmt.Print("|")
		for x := 0; x < tetris.width; x++ {
			node := Node{x, y}
			_, falling := shapeNodes[node]
			_, rest := tetris.chamber[node]
			if falling {
				fmt.Print("@")
			} else if rest {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|\n")
	}
	fmt.Print("+")
	for i := 0; i < tetris.width; i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n\n")
}

// Plays Tetris for the defined number of pieces. Returns total tower height
func (tetris *Tetris) Go(num int) int {
	drawStart := -1
	drawSteps := -1
	for n := 0; n < num; n++ {
		tetris.dropBlock(drawStart >= 0 && drawStart <= n, drawSteps == n)
	}
	return tetris.currentHeight
}

// Drops one block, optionally drawing all drop steps or starting state
func (tetris *Tetris) dropBlock(drawStart bool, drawSteps bool) {
	shape := tetris.getShape()
	if drawStart {
		fmt.Printf("Begin, current height = %v\n", tetris.currentHeight)
		tetris.draw(shape)
	}
	for {
		jet := tetris.getJet()
		tetris.move(shape, jet)
		if drawSteps {
			fmt.Printf("Pushes %v:\n", string(jet))
			tetris.draw(shape)
		}
		if tetris.drop(shape) {
			for node := range shape.getChamberNodes() {
				tetris.chamber[node] = struct{}{}
			}
			blockHeight := shape.position[1] + shape.height - 1
			if blockHeight > tetris.currentHeight {
				tetris.currentHeight = blockHeight
			}
			tetris.blocks++
			break
		}
		if drawSteps {
			fmt.Printf("Falls:\n")
			tetris.draw(shape)
		}
	}
	if drawSteps {
		emptyShape := Shape{nodes: []Node{}, position: Node{0, 0}, height: 0}
		fmt.Printf("Rest, current height = %v\n", tetris.currentHeight)
		tetris.draw(&emptyShape)
	}
}

type heightHash struct {
	blocks int
	height int
	hash   []int
}

// Drops blocks until a looping pattern is found, then returns the number of loop blocks and loop height
func (tetris *Tetris) findLoop() (int, int) {
	hashes := map[Node]heightHash{}
	for {
		tetris.dropBlock(false, false)
		// Compare top 30 blocks, and store with jet and shape index as "hash"
		hash := heightHash{blocks: int(tetris.blocks), height: int(tetris.currentHeight)}
		i := 0
		for y := tetris.currentHeight; y > tetris.currentHeight-30 && y > 0; y-- {
			for x := 0; x < tetris.width; x++ {
				value := 0
				if tetris.isFilled(Node{x, y}) {
					value = 1
				}
				hash.hash = append(hash.hash, value)
				i++
			}
		}
		idx := Node{tetris.shpIdx, tetris.jetIdx}
		last, seen := hashes[idx]
		if seen && reflect.DeepEqual(last.hash, hash.hash) {
			return tetris.blocks - last.blocks, tetris.currentHeight - last.height
		}
		hashes[idx] = hash
	}
}

// Drops blocks until a looping pattern is found, then returns the number of loop blocks and tower height
func (tetris *Tetris) findLoopUsingDelta() (int, int) {
	heightHash := map[Node][3]int{}
	for {
		tetris.dropBlock(false, false)
		// Compare delta height with last similar index, and store with jet and shape index as "hash"
		idx := Node{tetris.shpIdx, tetris.jetIdx}
		last, seen := heightHash[idx]
		if seen {
			last_dy := last[1] - last[0]
			current_dy := tetris.currentHeight - last[1]
			delta_blocks := tetris.blocks - last[2]
			if last_dy == current_dy {
				return delta_blocks, current_dy
			}
			heightHash[idx] = [3]int{last[1], int(tetris.currentHeight), delta_blocks}
		} else {
			heightHash[idx] = [3]int{0, int(tetris.currentHeight), tetris.blocks}
		}
	}
}

// Keeps track of shape properties, with grid node map and dimensions
type Shape struct {
	nodes    []Node
	height   int
	width    int
	position Node // Chamber index of bottom-left node
}

func NewShape(nodes []Node, width int, height int) *Shape {
	return &Shape{nodes: nodes, width: width, height: height}
}

// Get chamber position of shape nodes
func (shape *Shape) getChamberNodes() map[Node]struct{} {
	nodes := map[Node]struct{}{}
	for _, node := range shape.nodes {
		nodes[Node{node[0] + shape.position[0], node[1] + shape.position[1]}] = struct{}{}
	}
	return nodes
}

// Parses puzzle input from txt file.
// Returns tetris instace with shapes and jet pattern
func readInput(path string) *Tetris {
	inputBytes, err := os.ReadFile(path)
	check(err)
	inputString := string(inputBytes)
	jets := []rune{}
	for _, r := range inputString {
		jets = append(jets, r)
	}
	tetris := NewTetris(jets, 7)
	return tetris
}

// Part 1: How many units tall will the tower of rocks be after 2022 rocks have stopped falling?
func solvePart1(tetris *Tetris) int {
	return tetris.Go(2022)
}

// Part 2: How tall will the tower be after 1000000000000 rocks have stopped?
func solvePart2(tetris *Tetris) int {
	totalBlocks := 1000000000000
	loopBlocks, loopHeight := tetris.findLoop()

	remainingBlocks := totalBlocks - tetris.blocks
	topBlocks := remainingBlocks % loopBlocks
	heightWithoutLoops := tetris.Go(topBlocks)

	remainingBlocks = totalBlocks - tetris.blocks
	loops := int(remainingBlocks / loopBlocks)
	totalLoopHeight := loops * loopHeight

	return heightWithoutLoops + totalLoopHeight
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
	log.Printf("Tower height after 2022 rocks: %v\n", answer1)
	log.Printf("Tower height after 1000000000000 rocks: %v\n", answer2)
}
