// Advent of code 2022, day 7
// https://adventofcode.com/2022/day/7
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

// Keeps track of folder properties, subfolders, files, total file size and parent folder
type Folder struct {
	name       string
	parent     *Folder
	subfolders map[string]*Folder
	files      map[string]int
	totalSize  int
}

func NewFolder(name string, parent *Folder) *Folder {
	folder := Folder{name: name, parent: parent}
	folder.subfolders = make(map[string]*Folder)
	folder.files = make(map[string]int)
	return &folder
}

// Calculates total folder size with recursion through subfolders
func (folder *Folder) CalculateTotalSize() int {
	totalSize := 0
	for _, fileSize := range folder.files {
		totalSize += fileSize
	}
	for _, subfolder := range folder.subfolders {
		subfolder.CalculateTotalSize()
		totalSize += subfolder.totalSize
	}
	folder.totalSize = totalSize
	return totalSize
}

// Returns (sub)folder with total size closest to target size
func (folder *Folder) GetDeleteCandidate(targetSize int, currentCandidate *Folder) *Folder {
	for _, subfolder := range folder.subfolders {
		currentCandidate = subfolder.GetDeleteCandidate(targetSize, currentCandidate)
	}
	if folder.totalSize >= targetSize &&
		(currentCandidate == nil || currentCandidate.totalSize > folder.totalSize) {
		currentCandidate = folder
	}
	return currentCandidate
}

// Recursively sums folder sizes based on a given max total folder size
// Pass in a maxSize of -1 to ignore parameter
func (folder *Folder) GetTotalSize(maxSize int) (sumTotalSize int) {
	sumTotalSize = 0
	if maxSize == -1 || folder.totalSize <= maxSize {
		sumTotalSize += folder.totalSize
	}
	for _, subfolder := range folder.subfolders {
		sumTotalSize += subfolder.GetTotalSize(maxSize)
	}
	return sumTotalSize
}

// Parses puzzle input from txt file.
// Returns root folder complete with size calculation of content.
func readInput(path string) (root *Folder) {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(string(inputBytes), "\n")

	root = NewFolder("/", nil)
	cwd := root
	for _, line := range lines {
		parts := strings.Fields(line)
		if parts[0] == "$" {
			if parts[1] == "cd" {
				// Change directory
				if parts[2] == ".." {
					cwd = cwd.parent
				} else if parts[2] == "/" {
					cwd = root
				} else {
					cwd = cwd.subfolders[parts[2]]
				}
			} else if parts[1] == "ls" {
				// Content is processed in next parts
				continue
			} else {
				panic("unknown $ command: " + parts[1])
			}
		} else if parts[0] == "dir" {
			// New subfolder
			subfolder := NewFolder(parts[1], cwd)
			cwd.subfolders[parts[1]] = subfolder
		} else {
			// New file
			cwd.files[parts[1]], err = strconv.Atoi(parts[0])
			check(err)
		}
	}
	// Perform initial folder size calculations
	root.CalculateTotalSize()
	return root
}

// Part 1: What is the sum of the total sizes of directories with a total size of at most 100000?
func solvePart1(root *Folder) int {
	sizeLimit := 100000
	return root.GetTotalSize(sizeLimit)
}

// Part 2: What is the total size of best deletion candidate?
func solvePart2(root *Folder) int {
	totalSpace := 70000000
	unusedSpace := totalSpace - root.totalSize
	targetSize := 30000000 - unusedSpace
	var currentCandidate *Folder
	return root.GetDeleteCandidate(targetSize, currentCandidate).totalSize
}

// Solves puzzle parts. Split up for benchmarking
func solvePuzzle() (int, int) {
	root := readInput(inputPath)
	answer1 := solvePart1(root)
	answer2 := solvePart2(root)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Sum of total sizes of directories at most 100000: %v\n", answer1)
	log.Printf("Total size of best folder deletion candidate: %v\n", answer2)
}
