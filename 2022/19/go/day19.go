// Advent of code 2022, day 19
// https://adventofcode.com/2022/day/19
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

// Costs type. Map of robot types pointing to array of cost: [ore, clay, obsidian]
type Costs map[int][3]int // (0 = ore, 1 = clay, 2 = obsidian, 3 = geode)

// Keeps track of blueprints, with costs for robots and eventually max geodes cracked
type Blueprint struct {
	id int
	// Robot types pointing to array of cost: [ore, clay, obsidian]
	costs     Costs // (0 = ore, 1 = clay, 2 = obsidian, 3 = geode)
	maxCosts  [3]int
	maxGeodes int
	quality   int
}

func NewBlueprint(id int, costs Costs) *Blueprint {
	blueprint := Blueprint{id: id, costs: costs, maxCosts: [3]int{0, 0, 0}}
	for _, cost := range costs {
		if cost[0] > blueprint.maxCosts[0] {
			blueprint.maxCosts[0] = cost[0]
		}
		if cost[1] > blueprint.maxCosts[1] {
			blueprint.maxCosts[1] = cost[1]
		}
		if cost[2] > blueprint.maxCosts[2] {
			blueprint.maxCosts[2] = cost[2]
		}
	}
	return &blueprint
}

// Runs blueprint simulation using DFS to find max geodes cracked
// Returns blueprint quality level
func (blueprint *Blueprint) Optimize(ticks int) int {
	queue := []*Simulation{NewSimulation(blueprint)}
	for {
		// Break if queue empty
		if len(queue) == 0 {
			break
		}
		// Pop last queue item
		sim := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		// Check if time's up
		if sim.ticks == ticks {
			if sim.geodes > blueprint.maxGeodes {
				blueprint.maxGeodes = sim.geodes
			}
			continue
		}

		// Kill timeline if unable to beat current best outcome even if only producing geode robots
		remaining := ticks - sim.ticks
		if sim.geodes+sim.robots[3]*remaining+potentialGeodes(remaining) < blueprint.maxGeodes {
			continue
		}

		// If you can, always build a geode cracking robot and ignore the rest
		if sim.ore >= blueprint.costs[3][0] && sim.obsidian >= blueprint.costs[3][2] {
			next := CopySimulation(sim)
			next.ore -= blueprint.costs[3][0]
			next.obsidian -= blueprint.costs[3][2]
			next.tick(1)
			next.robots[3]++
			queue = append(queue, next)
			continue
		}
		// Split out into robot options
		for id, count := range sim.robots {
			// Don't produce robot if count same as largest resource cost
			if id < 3 && count >= blueprint.maxCosts[id] {
				continue
			}
			// Don't produce if resource unavailable
			if blueprint.costs[id][1] > 0 && sim.robots[1] == 0 || blueprint.costs[id][2] > 0 && sim.robots[2] == 0 {
				continue
			}

			next := CopySimulation(sim)
			// Consumes resources to build robot. Handles timeline by ticking until resources available
			costs := next.blueprint.costs[id]
			for {
				if next.ticks == ticks || (next.ore >= costs[0] && next.clay >= costs[1] && next.obsidian >= costs[2]) {
					break
				}
				next.tick(1)
			}
			// Check if time's up
			if next.ticks == ticks {
				if next.geodes > blueprint.maxGeodes {
					blueprint.maxGeodes = next.geodes
				}
				continue
			}
			// Start building
			next.ore -= costs[0]
			next.clay -= costs[1]
			next.obsidian -= costs[2]
			// Collect resources
			next.tick(1)
			// Increment robot count
			next.robots[id]++
			// Append to simulation queue
			queue = append(queue, next)
		}
	}
	// Calculate blueprint quality
	blueprint.quality = blueprint.id * blueprint.maxGeodes
	return blueprint.quality
}

// Calculates potential geodes if only producing geode crackers for remaining ticks
func potentialGeodes(remaining int) int {
	geodes := remaining - 1
	if remaining > 0 {
		geodes += potentialGeodes(remaining - 1)
	}
	return geodes
}

// Keeps track of blueprint simulation parameters
type Simulation struct {
	blueprint *Blueprint
	ticks     int
	robots    map[int]int // 0 = ore, 1 = clay, 2 = obsidian, 3 = geode
	ore       int
	clay      int
	obsidian  int
	geodes    int
}

func NewSimulation(blueprint *Blueprint) *Simulation {
	simulation := Simulation{
		blueprint: blueprint,
		robots:    map[int]int{0: 1, 1: 0, 2: 0, 3: 0},
	}
	return &simulation
}

func CopySimulation(prototype *Simulation) *Simulation {
	simulation := Simulation{
		blueprint: prototype.blueprint, ticks: prototype.ticks,
		ore: prototype.ore, clay: prototype.clay,
		obsidian: prototype.obsidian, geodes: prototype.geodes,
	}
	simulation.robots = map[int]int{}
	for id, count := range prototype.robots {
		simulation.robots[id] = count
	}
	return &simulation
}

// Advances simulation time by harvesting resources
func (sim *Simulation) tick(ticks int) {
	sim.ore += sim.robots[0] * ticks
	sim.clay += sim.robots[1] * ticks
	sim.obsidian += sim.robots[2] * ticks
	sim.geodes += sim.robots[3] * ticks
	sim.ticks += ticks
}

// Parses puzzle input from txt file.
// Returns slice of blueprint instances
func readInput(path string) []*Blueprint {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")
	blueprints := []*Blueprint{}
	for i, line := range lines {
		costs := Costs{}
		split := strings.Split(line, ".")

		// Ore robot
		oreRobotFields := strings.Fields(split[0])
		oreRobotOreCost := strToInt(oreRobotFields[len(oreRobotFields)-2])
		costs[0] = [3]int{oreRobotOreCost, 0, 0}
		// Clay robot
		clayRobotFields := strings.Fields(split[1])
		clayRobotOreCost := strToInt(clayRobotFields[len(clayRobotFields)-2])
		costs[1] = [3]int{clayRobotOreCost, 0, 0}
		// Obsidian robot
		obsidianRobotFields := strings.Fields(split[2])
		obsidianRobotOreCost := strToInt(obsidianRobotFields[len(obsidianRobotFields)-5])
		obsidianRobotClayCost := strToInt(obsidianRobotFields[len(obsidianRobotFields)-2])
		costs[2] = [3]int{obsidianRobotOreCost, obsidianRobotClayCost, 0}
		// Geode robot
		geodeRobotFields := strings.Fields(split[3])
		geodeRobotOreCost := strToInt(geodeRobotFields[len(geodeRobotFields)-5])
		geodeRobotObsidianCost := strToInt(geodeRobotFields[len(geodeRobotFields)-2])
		costs[3] = [3]int{geodeRobotOreCost, 0, geodeRobotObsidianCost}

		blueprint := NewBlueprint(i+1, costs)
		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}

func strToInt(s string) int {
	value, err := strconv.Atoi(s)
	check(err)
	return value
}

// Part 1: What do you get if you add up the quality level of all of the blueprints in your list?
func solvePart1(blueprints []*Blueprint) int {
	sumQuality := 0
	for _, blueprint := range blueprints {
		blueprint.Optimize(24)
		sumQuality += blueprint.quality
	}
	return sumQuality
}

// Part 2: What do you get if you multiply max geodes together?
func solvePart2(blueprints []*Blueprint) int {
	answer := 0
	for i, blueprint := range blueprints {
		blueprint.Optimize(32)
		if answer == 0 {
			answer = blueprint.maxGeodes
		} else {
			answer *= blueprint.maxGeodes
		}
		if i >= 2 {
			break
		}
	}
	return answer
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
	log.Printf("Quality level of all blueprints: %v\n", answer1)
	log.Printf("Top three blueprint geodes multiplied: %v\n", answer2)
}
