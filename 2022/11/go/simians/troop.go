package simians

// A group of monkeys is called a troop.
// Keeps track of monkeys and round order
type Troop struct {
	Monkeys            []*Monkey
	round              int
	WorryReduces       bool
	productOfDivisible int
}

// Inits new rope object with moves
func NewTroop(monkeys []*Monkey) *Troop {
	troop := Troop{Monkeys: monkeys, WorryReduces: true}
	troop.productOfDivisible = 1
	for _, monkey := range monkeys {
		troop.productOfDivisible *= monkey.testValue
	}
	return &troop
}

// Runs n rounds of monkey business
func (troop *Troop) Go(numRounds int) {
	for r := 0; r < numRounds; r++ {
		for _, monkey := range troop.Monkeys {
			for _, item := range monkey.items {
				value := monkey.operation(item)
				if troop.WorryReduces {
					value = int(monkey.operation(item) / 3)
				} else {
					// Well shit..
					value = value % troop.productOfDivisible
				}
				var target *Monkey
				if monkey.test(value) {
					target = troop.Monkeys[monkey.onTrueMonkey]
				} else {
					target = troop.Monkeys[monkey.onFalseMonkey]
				}
				target.items = append(target.items, value)
				monkey.InspectionCount++
			}
			monkey.items = []int{}
		}
		troop.round++
	}
}
