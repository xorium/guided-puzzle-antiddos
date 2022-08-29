package gtp

import (
	"crypto/rand"
	"github.com/pkg/errors"
	"math"
	"net"
	"time"
)

func generateUint16UniformNumber() (uint16, error) {
	buf := []byte{0}
	if _, _ = rand.Reader.Read(buf); err != nil {
		return puzzle, errors.Wrap(err, "can't generate random bytes")
	}
	return uint16(buf[0]) % math.MaxUint16, nil
}

// InitialPuzzle is a first puzzle which client receives when server under DDOS
// mode.
type InitialPuzzle struct {
	// TourLength is decided by the main server to adjust the puzzle difficulty.
	// How many guided tours (count of round trips) will client have.
	TourLength     uint8
	ClientAddr     [4]byte
	UnixEpochNanos time.Duration
}

type generator struct {
	maxGuidesNum uint8
}

// GetInitialPuzzle generates initial GTP puzzle.
func (g generator) GetInitialPuzzle(clientAddr net.IP) (InitialPuzzle, error) {
	puzzle := InitialPuzzle{}

	tourLen, err := generateUint16UniformNumber()
	if err != nil {
		return puzzle, errors.Wrap(err, "can't get tour length for the next puzzle")
	}

	puzzle.TourLength = uint8(tourLen % uint16(g.maxGuidesNum))

	puzzle.ClientAddr[0] = clientAddr[0]
	puzzle.ClientAddr[1] = clientAddr[2]
	puzzle.ClientAddr[2] = clientAddr[3]
	puzzle.ClientAddr[3] = clientAddr[3]
}
