// Advent of Code, day 2
// https://adventofcode.com/2022/day/2
package main

import (
	"fmt"
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

// Super for shapes, subclassed into rock, paper, scissors etc.
type Shape struct {
	ids    []string
	name   string
	points int
	beats  map[string]struct{}
}

// Shape constructor
func newShape(ids []string, name string, points int) *Shape {
	shape := Shape{ids: ids, name: name, points: points}
	shape.beats = make(map[string]struct{})
	return &shape
}

// Handles round simulation and score calculation
type Round struct {
	playerShape   Shape
	opponentShape Shape
	outcome       string // "win"/"draw"/"loss"
	score         int
}

// Constructor for rounds. Calculates outcome and points based on shapes
func newRound(playerShape Shape, opponentShape Shape) *Round {
	round := Round{playerShape: playerShape, opponentShape: opponentShape}
	round.outcome = round.getOutcome()
	round.score = round.calculateScore()
	return &round
}

// Calculates outcome based on player and opponent shapes
func (round Round) getOutcome() string {
	if _, ok := round.playerShape.beats[round.opponentShape.name]; ok {
		return "win"
	}
	if _, ok := round.opponentShape.beats[round.playerShape.name]; ok {
		return "loss"
	}
	return "draw"
}

// Calculates round score based on outcome and player shape
func (round Round) calculateScore() (score int) {
	score = round.playerShape.points
	if round.outcome == "win" {
		score += 6
	} else if round.outcome == "draw" {
		score += 3
	}
	return score
}

// Parses puzzle input from txt file.
// Returns a list of strings representing shapes played
func readInput(inputPath string) (rounds []string) {
	inputBytes, err := os.ReadFile(inputPath)
	check(err)
	input := string(inputBytes)

	for _, line := range strings.Split(input, "\n") {
		rounds = append(rounds, strings.TrimSpace(line))
	}
	return rounds
}

// Inits shape types and rules. Returns a map of shape names pointing to shape structs
// Needs rework to init based on some kind of config if shapes unknown
func initShapes() (shapes map[string]*Shape) {
	shapes = make(map[string]*Shape)

	rockIds := []string{"A", "X"}
	rock := newShape(rockIds, "rock", 1)
	paperIds := []string{"B", "Y"}
	paper := newShape(paperIds, "paper", 2)
	scissorIds := []string{"C", "Z"}
	scissors := newShape(scissorIds, "scissors", 3)

	// The holy trinity
	rock.beats[scissors.name] = struct{}{}
	paper.beats[rock.name] = struct{}{}
	scissors.beats[paper.name] = struct{}{}

	shapes[rock.name] = rock
	shapes[paper.name] = paper
	shapes[scissors.name] = scissors
	return shapes
}

// Simulates a series of rounds and returns the score
func simulateRounds(rounds []string, shapes map[string]*Shape) (totalScore int) {
	totalScore = 0
	for _, roundIds := range rounds {
		shapeIds := strings.Split(roundIds, " ")
		opponentShape, err := getShape(shapeIds[0], shapes)
		check(err)
		playerShape, err := getShape(shapeIds[1], shapes)
		check(err)
		round := newRound(*playerShape, *opponentShape)
		totalScore += round.score
	}
	return totalScore
}

// Gets shape from id and available shapes
func getShape(id string, shapes map[string]*Shape) (*Shape, error) {
	for _, shape := range shapes {
		// I think I might hate Go
		for _, shapeId := range shape.ids {
			if id == shapeId {
				return shape, nil
			}
		}
	}
	return &Shape{}, fmt.Errorf("invalid shape id: " + id)
}

// Simulates a series of rounds where we choose the winning strategy, then returns the score
func executeStrategy(rounds []string, shapes map[string]*Shape) (totalScore int) {
	totalScore = 0
	var playerShape *Shape
	for _, roundIds := range rounds {
		shapeIds := strings.Split(roundIds, " ")
		opponentShape, err := getShape(shapeIds[0], shapes)
		check(err)
		playerOutcome := shapeIds[1]

		// Too late for this shit, hardcoding
		if playerOutcome == "X" {
			if opponentShape.name == "rock" {
				playerShape = shapes["scissors"]
			} else if opponentShape.name == "paper" {
				playerShape = shapes["rock"]
			} else {
				playerShape = shapes["paper"]
			}
		} else if playerOutcome == "Y" {
			playerShape = opponentShape
		} else {
			if opponentShape.name == "rock" {
				playerShape = shapes["paper"]
			} else if opponentShape.name == "paper" {
				playerShape = shapes["scissors"]
			} else {
				playerShape = shapes["rock"]
			}
		}

		round := newRound(*playerShape, *opponentShape)
		totalScore += round.score
	}
	return totalScore
}

// Parses input as txt into list of rounds, then simulates points
func main() {
	rounds := readInput(inputPath)
	shapes := initShapes()

	// Part 1
	points := simulateRounds(rounds, shapes)
	fmt.Println(points)

	// Part 2
	points = executeStrategy(rounds, shapes)
	fmt.Println(points)
}
