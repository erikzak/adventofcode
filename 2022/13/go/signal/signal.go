package signal

// Keeps track of packets received as slices of packet pairs.
// Has methods for checking packet ordering.
type Signal struct {
	PacketPairs [][]Packet // Pairs of packets
}

func NewSignal(packetPairs [][]Packet) (signal Signal) {
	return Signal{PacketPairs: packetPairs}
}

// Checks packet ordering, and returns the sum of correctly ordered packet indexes (+1)
func (signal Signal) CheckOrdering() (sumCorrectIdx int) {
	for i, pair := range signal.PacketPairs {
		if pair[0].IsOrdered(pair[1]) {
			sumCorrectIdx += i + 1
		}
	}
	return sumCorrectIdx
}
