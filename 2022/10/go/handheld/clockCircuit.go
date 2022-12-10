package handheld

import (
	"log"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Keeps track of clock circuit register and sum of interesting cycles
type ClockCircuit struct {
	register           int
	cycle              int
	interestingCycles  []int
	sumSignalStrengths int
}

// Inits new clock circuit object with moves
func NewClockCircuit() *ClockCircuit {
	clockCircuit := ClockCircuit{}
	return &clockCircuit
}

// Advances cycle by one tick, checking for interesting signals
func (clockCircuit *ClockCircuit) tick() {
	clockCircuit.cycle++
	clockCircuit.trackInterestingSignals()
}

// Checks if signal strengt of current cycle should be tracked
func (clockCircuit *ClockCircuit) trackInterestingSignals() {
	for i := range clockCircuit.interestingCycles {
		if clockCircuit.interestingCycles[i] == clockCircuit.cycle {
			signalStrength := clockCircuit.cycle * clockCircuit.register
			clockCircuit.sumSignalStrengths += signalStrength
			break
		}
	}
}
