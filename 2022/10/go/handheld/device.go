package handheld

import (
	"strconv"
	"strings"
)

// Handheld device. Keeps track of clock circuit and screen.
// Implements methods for executing instructions for modifying clock circuit
// register and screen updates
type Device struct {
	clockCircuit *ClockCircuit
	screen       *Screen
	instructions []string
}

func NewDevice(instructions []string, screenHeight int, screenWidth int) *Device {
	device := Device{instructions: instructions}
	device.clockCircuit = NewClockCircuit()
	device.screen = NewScreen(screenWidth, screenHeight)
	return &device
}

// Resets clock circuit register and cycles
func (device *Device) reset(interestingCycles []int) {
	device.clockCircuit.cycle = 1
	device.clockCircuit.interestingCycles = interestingCycles
	device.clockCircuit.register = 1
	device.clockCircuit.sumSignalStrengths = 0
}

// Executes instructions
func (device *Device) ExecuteInstructions(interestingCycles []int) {
	device.reset(interestingCycles)
	for _, command := range device.instructions {
		if command == "noop" {
			device.noop()
		} else {
			value, err := strconv.Atoi(strings.Split(command, " ")[1])
			check(err)
			device.addx(value)
		}
	}
}

// noop command: takes one cycle to complete, no other effect
func (device *Device) noop() {
	device.screen.drawPixel(device.clockCircuit.cycle, device.clockCircuit.register)
	device.clockCircuit.tick()
}

// addx command: takes two cycle to complete, adds value to register
func (device *Device) addx(value int) {
	device.screen.drawPixel(device.clockCircuit.cycle, device.clockCircuit.register)
	device.clockCircuit.tick()
	device.screen.drawPixel(device.clockCircuit.cycle, device.clockCircuit.register)
	device.clockCircuit.register += value
	device.clockCircuit.tick()
}

// Returns current clock circuit sum of interesting signal strengths
func (device *Device) GetSumSignalStrengths() int {
	return device.clockCircuit.sumSignalStrengths
}

// Returns rendered screen image
func (device *Device) GetImage() []string {
	return device.screen.getImage()
}
