// Advent of code 2022, day 9
// https://adventofcode.com/2022/day/9
package main

import (
	"fmt"
	"log"
	"math"
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

// ---------------------------------------------------------------------------
// Keeps track of rope knots, list of moves and locations visited.
// Implements methods for moving head and knots.
type Rope struct {
	knots           []*Knot
	moves           []string
	tailLocsVisited map[string]struct{}
}

// Inits new rope object with moves
func NewRope(moves []string) *Rope {
	rope := Rope{moves: moves}
	return &rope
}

// Resets head and knot locations, slice of knots and map of visited locations
func (rope *Rope) reset(length int) {
	rope.knots = []*Knot{}
	for i := 0; i < length; i++ {
		rope.knots = append(rope.knots, NewKnot())
	}
	rope.knots[length-1].isTail = true
	rope.tailLocsVisited = map[string]struct{}{}
	rope.addTailLoc()
}

// Executes moveset with given rope length. Returns number of locations visited by tail
func (rope *Rope) executeMoves(length int) int {
	rope.reset(length)
	for _, move := range rope.moves {
		instructions := strings.Split(move, " ")
		direction := instructions[0]
		length, err := strconv.Atoi(instructions[1])
		check(err)
		for step := 0; step < length; step++ {
			rope.moveHead(direction)
			for i := 1; i < len(rope.knots); i++ {
				rope.follow(rope.knots[i], *rope.knots[i-1])
			}
		}
	}
	return len(rope.tailLocsVisited)
}

// Moves head in the specified direction
func (rope *Rope) moveHead(direction string) {
	head := rope.knots[0]
	if direction == "U" {
		head.y++
	} else if direction == "D" {
		head.y--
	} else if direction == "L" {
		head.x--
	} else if direction == "R" {
		head.x++
	}
}

// Moves a rope knot incrementally closer to its leading knot if distance is more than one location.
// Also keeps track of unique set of visitied locations
func (rope *Rope) follow(knot *Knot, target Knot) {
	if knot.x == target.x && knot.y == target.y {
		return
	}
	dx, dy := knot.getDist(target)
	for distanceBetween(dx, dy) >= 2 {
		if dx > 0 {
			knot.x++
		} else if dx < 0 {
			knot.x--
		}
		if dy > 0 {
			knot.y++
		} else if dy < 0 {
			knot.y--
		}
		if knot.isTail {
			rope.addTailLoc()
		}
		dx, dy = knot.getDist(target)
	}
}

// Adds current rope tail to map of visited locations
func (rope *Rope) addTailLoc() {
	tail := rope.knots[len(rope.knots)-1]
	key := fmt.Sprint(tail.x) + ":" + fmt.Sprint(tail.y)
	rope.tailLocsVisited[key] = struct{}{}
}

// Calculates distance between two points using dx and dy
func distanceBetween(dx, dy int) float64 {
	return math.Sqrt(float64(dx*dx + dy*dy))
}

// ---------------------------------------------------------------------------
// Keeps track of knot location
type Knot struct {
	x      int
	y      int
	isTail bool
}

func NewKnot() *Knot {
	knot := Knot{x: 0, y: 0}
	return &knot
}

// Returns the X and distance between knot and target
func (knot *Knot) getDist(target Knot) (dx, dy int) {
	dx = target.x - knot.x
	dy = target.y - knot.y
	return dx, dy
}

// ---------------------------------------------------------------------------
// Parses puzzle input from txt file.
// Returns rope instance with list of moves
func readInput(path string) (rope *Rope) {
	inputBytes, err := os.ReadFile(path)
	check(err)
	moves := strings.Split(string(inputBytes), "\n")
	rope = NewRope(moves)
	return rope
}

// Part 1: How many positions does the tail of the rope(2) visit at least once?
func solvePart1(rope *Rope) (tailLocsVisited int) {
	ropeLength := 2
	tailLocsVisited = rope.executeMoves(ropeLength)
	return tailLocsVisited
}

// Part 2: How many positions does the tail of the rope(10) visit at least once?
func solvePart2(rope *Rope) (tailLocsVisited int) {
	ropeLength := 10
	tailLocsVisited = rope.executeMoves(ropeLength)
	return tailLocsVisited
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
	log.Printf("Positions visited with rope length 2: %v\n", answer1)
	log.Printf("Positions visited with rope length 10: %v\n", answer2)
}
