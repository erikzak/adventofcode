package crane

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Crane structure. Keeps track of stacks and planned moves.
// Moves crates between stacks.
type Crane struct {
	stacks       map[rune]Stack
	plannedMoves []string
}

// Crane constructor. Defines initial stack configuration
func NewCrane(stacks map[rune]Stack, plannedMoves []string) Crane {
	crane := Crane{stacks: stacks, plannedMoves: plannedMoves}
	return crane
}

func (crane *Crane) ExecuteMoves(singleCrate bool) {
	for _, move := range crane.plannedMoves {
		fields := strings.Fields(move)
		count, err := strconv.Atoi(fields[1])
		check(err)
		fromStackId := rune(fields[3][0])
		toStackId := rune(fields[5][0])
		crane.MoveCrates(count, fromStackId, toStackId, singleCrate)
	}
}

// Moves crates between stacks based on planned moves
func (crane *Crane) MoveCrates(count int, fromStackId rune, toStackId rune, singleCrate bool) {
	fromStack, toStack := crane.stacks[fromStackId], crane.stacks[toStackId]
	movedCrates := append([]rune(nil), fromStack.crates[:count]...)
	if singleCrate {
		reverse(movedCrates)
	}
	toStack.crates = append(movedCrates, toStack.crates...)
	fromStack.crates = fromStack.crates[count:]
	crane.stacks[fromStackId] = fromStack
	crane.stacks[toStackId] = toStack
}

// Returns string of crates on top of stacks
func (crane *Crane) GetTopCrates() string {
	topCrates := ""
	// Iterate over map items sorted by id
	ids := make([]rune, 0, len(crane.stacks))
	for id := range crane.stacks {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	for _, id := range ids {
		topCrates += string(crane.stacks[id].crates[0])
	}
	return topCrates
}

// Reverses slide order
func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
