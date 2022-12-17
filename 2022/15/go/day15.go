// Advent of code 2022, day 15
// https://adventofcode.com/2022/day/15
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

// Define node map unit as XY pair
type Node [2]int

// Define sensor struct to keep track of node, beacon and Manhattan distance between them
type Sensor struct {
	node             Node
	beacon           Node
	distanceToBeacon int
}

// Returns all nodes at sensor boundary + offset
func (sensor Sensor) getBoundary(offset int) (nodes []Node) {
	distance := sensor.distanceToBeacon + offset
	corners := []Node{
		{sensor.node[0], sensor.node[1] + distance},
		{sensor.node[0] + distance, sensor.node[1]},
		{sensor.node[0], sensor.node[1] - distance},
		{sensor.node[0] - distance, sensor.node[1]},
	}
	for i := 0; i < len(corners)-1; i++ {
		from, to := corners[i], corners[i+1]
		nodes = append(nodes, from)
		dx, dy := to[0]-from[0], to[1]-from[1]
		incX, incY := 1, 1
		if dx < 0 {
			incX = -1
		}
		if dy < 0 {
			incY = -1
		}
		node := Node{from[0] + incX, from[1] + incY}
		for {
			nodes = append(nodes, node)
			if node == to {
				break
			}
			node = Node{node[0] + incX, node[1] + incY}
		}
	}
	return nodes
}

// Inits new sensor with location and beacon, calculating the Manhattan distance between them
func NewSensor(sensorNode Node, beaconNode Node) Sensor {
	sensor := Sensor{node: sensorNode, beacon: beaconNode}
	sensor.distanceToBeacon = calculateManhattanDistance(sensorNode, beaconNode)
	return sensor
}

// Returns manhattan distance between start and end node
func calculateManhattanDistance(start Node, end Node) int {
	return Abs(end[0]-start[0]) + Abs(end[1]-start[1])
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Cavern 2d vertical slice. Node map points to int specifying if node is
// blocked or not, and what is blocking
type Cavern struct {
	nodes             map[Node]int // -1 = unknown, 0 = beacon void, 1 = sensor, 2 = beacon
	sensors           map[Node]Sensor
	dim               [4]int // [xMin, xMax, zMin, zMax]
	maxBeaconDistance int
}

// Builds node map from sensors and beacons.
func NewCavern(sensors map[Node]Sensor) *Cavern {
	// Keep track of total dimensions: [xMin, xMax, zMin, zMax]
	dim := [4]int{}
	nodes := map[Node]int{}
	// Add sensors and beacons to node map
	maxBeaconDistance := 0
	for node, sensor := range sensors {
		nodes[node] = 1
		nodes[sensor.beacon] = 2
		dim = updateDimensions(dim, node)
		dim = updateDimensions(dim, sensor.beacon)
		if sensor.distanceToBeacon > maxBeaconDistance {
			maxBeaconDistance = sensor.distanceToBeacon
		}
	}
	cavern := Cavern{nodes: nodes, sensors: sensors, dim: dim, maxBeaconDistance: maxBeaconDistance}
	return &cavern
}

// Returns content of target node
func (cavern *Cavern) GetContent(node Node) int {
	// Check map for content
	content, exists := cavern.nodes[node]
	if !exists {
		content = -1
	}
	return content
}

// Checks if node is beacon void by calculating distance to all sensors and
// checking if the distance from the node to the sensor is less than or equal
// to the distance from sensor to its beacon
func (cavern *Cavern) IsVoid(node Node) bool {
	content := cavern.GetContent(node)
	if content == 2 {
		return false
	}
	if content == 0 || content == 1 {
		return true
	}
	for _, sensor := range cavern.sensors {
		distanceToSensor := calculateManhattanDistance(node, sensor.node)
		if distanceToSensor <= sensor.distanceToBeacon {
			return true
		}
	}
	return false
}

// Update node map dimensions based on given node
func updateDimensions(dim [4]int, node Node) [4]int {
	x, y := node[0], node[1]
	if dim[0] > x {
		dim[0] = x
	} else if dim[1] < x {
		dim[1] = x
	}
	if dim[2] > y {
		dim[2] = y
	} else if dim[3] < y {
		dim[3] = y
	}
	return dim
}

// Prints cavern map to console
func (cavern *Cavern) printMap() {
	charMap := map[int]string{
		-1: ".",
		0:  "#",
		1:  "S",
		2:  "B",
	}
	for y := cavern.dim[2]; y <= cavern.dim[3]; y++ {
		for x := cavern.dim[0]; x <= cavern.dim[1]; x++ {
			sign := " "
			content, ok := cavern.nodes[Node{x, y}]
			if ok {
				sign = charMap[content]
			}
			fmt.Print(sign)
		}
		fmt.Print("\n")
	}
}

// Parses puzzle input from txt file.
// Returns node map of cavern.
func readInput(path string) *Cavern {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")
	sensors := map[Node]Sensor{}
	// Generate sensor/beacon map pointing to signals from input strings
	for _, line := range lines {
		split := strings.Split(line, "=")
		sensorX, _ := strconv.Atoi(strings.Split(split[1], ",")[0])
		sensorY, _ := strconv.Atoi(strings.Split(split[2], ":")[0])
		beaconX, _ := strconv.Atoi(strings.Split(split[3], ",")[0])
		beaconY, _ := strconv.Atoi(strings.Split(split[4], ":")[0])
		sensorNode := Node{sensorX, sensorY}
		beaconNode := Node{beaconX, beaconY}
		sensor := NewSensor(sensorNode, beaconNode)
		sensors[sensorNode] = sensor
	}
	cavern := NewCavern(sensors)
	return cavern
}

// Part 1: In the row where y=2000000, how many positions cannot contain a beacon?
func solvePart1(cavern *Cavern, test bool) int {
	// Count voids in row
	voidCount := 0
	y := 2000000
	if test {
		y = 10
	}
	startX := cavern.dim[0] - cavern.maxBeaconDistance
	endX := cavern.dim[1] + cavern.maxBeaconDistance
	for x := startX; x < endX; x++ {
		node := Node{x, y}
		if cavern.IsVoid(node) {
			cavern.dim = updateDimensions(cavern.dim, node)
			content := cavern.GetContent(Node{x, y})
			if content == -1 {
				cavern.nodes[node] = 0
			}
			voidCount++
		}
	}
	if test {
		cavern.printMap()
	}
	return voidCount
}

// Part 2: What is the distress beacon tuning frequency?
func solvePart2(cavern *Cavern, test bool) int {
	maxIdx := 4000000
	if test {
		maxIdx = 20
	}
	for _, sensor := range cavern.sensors {
		for _, boundaryNode := range sensor.getBoundary(1) {
			if boundaryNode[0] < 0 || boundaryNode[0] > maxIdx || boundaryNode[1] < 0 || boundaryNode[1] > maxIdx {
				continue
			}
			if !cavern.IsVoid(boundaryNode) {
				content := cavern.GetContent(boundaryNode)
				if content == -1 {
					return boundaryNode[0]*4000000 + boundaryNode[1]
				}
			}
		}
	}
	panic("no distress beacon found")
}

// Solves puzzle parts. Split up for benchmarking
func solvePuzzle() (int, int) {
	input := readInput(inputPath)
	answer1 := solvePart1(input, false)
	answer2 := solvePart2(input, false)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Positions without beacons in row 2000000: %v\n", answer1)
	log.Printf("Distress beacon tuning frequency: %v\n", answer2)
}
