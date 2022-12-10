package crane

// Keeps track of stack number and crates.
// First crate slice item is bottom crate.
type Stack struct {
	id     rune
	crates []rune
}

// Stack constructor. Sets id and initial set of crates
func NewStack(id rune, crates []rune) Stack {
	return Stack{id: id, crates: crates}
}
