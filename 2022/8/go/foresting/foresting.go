package foresting

// Keeps track of forest grid, with rows and columns of trees
type Forest struct {
	rows         [][]*Tree
	columns      [][]*Tree
	visibleTrees int
}

func NewForest(rows [][]*Tree, columns [][]*Tree) *Forest {
	forest := Forest{rows: rows, columns: columns}
	for rowIdx := 0; rowIdx < len(forest.rows); rowIdx++ {
		for colIdx := 0; colIdx < len(forest.columns); colIdx++ {
			row, column := forest.rows[rowIdx], forest.columns[colIdx]
			tree := row[colIdx]
			tree.rowIdx, tree.colIdx = rowIdx, colIdx
			tree.row, tree.column = row, column
		}
	}
	return &forest
}

// Calculates tree visibility for all trees in the forest
func (forest *Forest) CalculateTreeVisibility() int {
	forest.visibleTrees = 0
	for _, row := range forest.rows {
		for _, tree := range row {
			tree.CalculateVisibility()
			if tree.isVisible {
				forest.visibleTrees++
			}
		}
	}
	return forest.visibleTrees
}

// Calculates scenic score for all trees in the forest
func (forest *Forest) CalculateTreeScenicScores() (highestScenicScore int) {
	for _, row := range forest.rows {
		for _, tree := range row {
			tree.CalculateScenicScore()
			if tree.scenicScore > highestScenicScore {
				highestScenicScore = tree.scenicScore
			}
		}
	}
	return highestScenicScore
}

// Keeps track of tree attributes, its row and column in the forest and calculated visibility
type Tree struct {
	height      int
	rowIdx      int
	colIdx      int
	row         []*Tree
	column      []*Tree
	isVisible   bool
	scenicScore int
}

func NewTree(height int) *Tree {
	tree := Tree{height: height}
	return &tree
}

// Calculates tree visibility based on position in grid and height of other trees
func (tree *Tree) CalculateVisibility() bool {
	// Visible if tree is on the edge of the grid
	if tree.rowIdx == 0 || tree.rowIdx == len(tree.row)-1 ||
		tree.colIdx == 0 || tree.colIdx == len(tree.column)-1 {
		tree.isVisible = true
		return tree.isVisible
	}

	tree.isVisible = false
	if tree.isVisibleInDirection(tree.row[:tree.colIdx]) ||
		tree.isVisibleInDirection(tree.row[tree.colIdx+1:]) ||
		tree.isVisibleInDirection(tree.column[:tree.rowIdx]) ||
		tree.isVisibleInDirection(tree.column[tree.rowIdx+1:]) {
		tree.isVisible = true
	}
	return tree.isVisible
}

// Calculates tree scenic score based on position in grid and visible trees
func (tree *Tree) CalculateScenicScore() int {
	tree.scenicScore = tree.getVisibleTreesInDirection(tree.row[:tree.colIdx], true) *
		tree.getVisibleTreesInDirection(tree.row[tree.colIdx+1:], false) *
		tree.getVisibleTreesInDirection(tree.column[:tree.rowIdx], true) *
		tree.getVisibleTreesInDirection(tree.column[tree.rowIdx+1:], false)
	return tree.scenicScore
}

// Checks if all given trees are lower than target height
func (tree *Tree) isVisibleInDirection(others []*Tree) bool {
	for _, other := range others {
		if other.height >= tree.height {
			return false
		}
	}
	return true
}

// getVisibleTrees
func (tree *Tree) getVisibleTreesInDirection(others []*Tree, reverse bool) (visibleTrees int) {
	if !reverse {
		for i := 0; i < len(others); i++ {
			visibleTrees++
			if others[i].height >= tree.height {
				break
			}
		}
	} else {
		for i := len(others) - 1; i > -1; i-- {
			visibleTrees++
			if others[i].height >= tree.height {
				break
			}
		}
	}
	return visibleTrees
}
