// Advent of code 2022, day 16
// https://adventofcode.com/2022/day/16
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

// Keeps track of valve system and cumulative pressure released
type Cave struct {
	valves            map[string]Valve
	workingValveCount int // Number of valves with actual flow rate
}

// Inits cave from map of valves
func NewCave(valves map[string]Valve) Cave {
	cave := Cave{valves: valves}
	for _, valve := range valves {
		// Keep track of working valves
		if valve.flowRate > 0 {
			cave.workingValveCount++
		}
		// Calculate distance to all other valves
		for name, target := range valves {
			if name == valve.name {
				continue
			}
			valve.distances[name] = cave.BFS(valve, target)
		}
	}
	return cave
}

// Uses breadth-first search to calculate distance between valves.
// Returns node distance as integer.
func (cave Cave) BFS(start Valve, end Valve) int {
	queue := []Valve{start}
	parents := map[string]string{}
	explored := map[string]struct{}{start.name: {}}
	for {
		node := queue[0]
		if node.name == end.name {
			return getBFSDistance(end.name, parents)
		}
		queue = queue[1:]
		for _, tunnel := range node.tunnels {
			_, done := explored[tunnel]
			if done {
				continue
			}
			explored[tunnel] = struct{}{}
			parents[tunnel] = node.name
			queue = append(queue, cave.valves[tunnel])
		}
	}
}

// Recreates path from start to end. Returns length of path
func getBFSDistance(end string, parents map[string]string) (distance int) {
	parent := parents[end]
	var found bool
	distance++
	for {
		parent, found = parents[parent]
		if !found {
			break
		}
		distance++
	}
	return distance
}

// Simulates possible routes and finds the one with most pressure released
func (cave Cave) getMaxPossiblePressureReleased(
	startingLocation string, maxMinutes int, workers int,
) int {
	root := NewRoute(cave.valves[startingLocation])
	bestRoute := root
	routes := []Route{root}
	for {
		if len(routes) == 0 {
			break
		}

		// Pop next queue item
		route := routes[0]
		routes = removeRoute(routes, 0)

		// Check if all valves opened in route
		if len(route.opened) == cave.workingValveCount {
			minutesLeft := maxMinutes - route.minutesElapsed
			for _, valve := range route.valves {
				route.pressureReleased += valve.flowRate * minutesLeft
			}
			if bestRoute.pressureReleased < route.pressureReleased {
				bestRoute = route
			}
			continue
		}

		current := root.start
		if len(route.valves) > 0 {
			current = route.valves[len(route.valves)-1]
		}
		for _, target := range cave.valves {
			// Check if valve has flow rate
			if target.flowRate == 0 {
				continue
			}
			// Check if valve has not already been opened on route
			_, opened := route.opened[target.name]
			if opened {
				continue
			}
			if (maxMinutes - route.minutesElapsed) < (current.distances[target.name] + 1) {
				// No time to get to next route destination.
				// Sum up remaining pressure and check result.
				finishedRoute := CopyRoute(route)
				minutesLeft := maxMinutes - finishedRoute.minutesElapsed
				for _, valve := range finishedRoute.valves {
					finishedRoute.pressureReleased += valve.flowRate * minutesLeft
				}
				if bestRoute.pressureReleased < finishedRoute.pressureReleased {
					bestRoute = finishedRoute
				}
				continue
			}
			// Simulate next route step
			nextStep := CopyRoute(route)
			minutesSpent := current.distances[target.name] + 1
			nextStep.minutesElapsed += minutesSpent
			// Accumulate released pressure
			for _, valve := range nextStep.valves {
				nextStep.pressureReleased += valve.flowRate * minutesSpent
			}
			// Keep track of route steps
			nextStep.valves = append(nextStep.valves, target)
			nextStep.opened[target.name] = struct{}{}
			routes = append(routes, nextStep)
		}
	}
	fmt.Print(bestRoute.history)
	return bestRoute.pressureReleased
}

func removeRoute(s []Route, i int) []Route {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Keeps track of a specific valve opening route
type Route struct {
	start            Valve
	valves           []Valve
	opened           map[string]struct{}
	minutesElapsed   int
	pressureReleased int
	history          string
}

func NewRoute(start Valve) Route {
	route := Route{
		start: start, valves: []Valve{}, opened: map[string]struct{}{},
		minutesElapsed: 0, pressureReleased: 0,
	}
	return route
}

func CopyRoute(route Route) Route {
	copy := Route{minutesElapsed: route.minutesElapsed, pressureReleased: route.pressureReleased, history: route.history}
	copy.valves = []Valve{}
	copy.valves = append(copy.valves, route.valves...)
	copy.opened = map[string]struct{}{}
	for name := range route.opened {
		copy.opened[name] = struct{}{}
	}
	return copy
}

func (route Route) getRouteSteps() string {
	steps := []string{}
	flowRate := 0
	for _, valve := range route.valves {
		steps = append(steps, valve.name)
		flowRate += valve.flowRate
	}
	valves := strings.Join(steps, ", ")
	return valves + fmt.Sprintf(" releasing %v pressure", flowRate)
}

// Keeps track of valve name, flow rate and paths to other valves
type Valve struct {
	name      string
	flowRate  int
	tunnels   []string
	distances map[string]int // Distances to other valves
}

// Inits new valve
func NewValve(name string, flowRate int, tunnels []string) Valve {
	valve := Valve{name: name, flowRate: flowRate, tunnels: tunnels, distances: map[string]int{}}
	return valve
}

// Parses puzzle input from txt file.
// Returns cave system of valves.
func readInput(path string) Cave {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")
	valves := map[string]Valve{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		name := split[1]
		flowRate, _ := strconv.Atoi(strings.Trim(strings.Split(split[4], "=")[1], ";"))
		tunnels := []string{}
		for _, tunnel := range split[9:] {
			tunnels = append(tunnels, strings.Trim(tunnel, ","))
		}
		valves[name] = NewValve(name, flowRate, tunnels)
	}
	cave := NewCave(valves)
	return cave
}

// Part 1: What is the most pressure you can release?
func solvePart1(cave Cave) int {
	mostReleasedPressure := cave.getMaxPossiblePressureReleased("AA", 30, 1)
	return mostReleasedPressure
}

// Part 2: With you and an elephant working together for 26 minutes, what is
// the most pressure you could release?
func solvePart2(cave Cave) int {
	mostReleasedPressure := cave.getMaxPossiblePressureReleased("AA", 26, 2)
	return mostReleasedPressure
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
	log.Printf("Most pressure that can be released: %v\n", answer1)
	log.Printf("Most pressure that can be released with elephant: %v\n", answer2)
}
