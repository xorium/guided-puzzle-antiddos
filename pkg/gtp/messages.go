package gtp

import (
	"net"
	"time"
)

type RegisteredRespMsg struct {
	ID           int
	SharedKey    string
	ActiveGuides []GuideInfoMsg
}

type GuideInfoMsg struct {
	ID         int
	SharedKey  string
	PublicAddr string
}

type ResourceRequestMsg struct{}

type ResourceResponseMsg struct {
	Quote string
}

// InitialPuzzleMsg is a first puzzle which crypto receives when server under
// anti-DDOS mode.
type InitialPuzzleMsg struct {
	// TourLength is decided by the main server to adjust the puzzle difficulty.
	// How many guided tours (count of round trips) will crypto have.
	TourLength uint16 `msgpack:"L"`
	// NextGuideAddr is the ID of the next tour gtp to follow.
	NextGuideID int `msgpack:"i"`
	// NextGuideAddr is the IP address of the next tour gtp to follow.
	NextGuideAddr net.IP `msgpack:"ip"`
	// UnixNanos coursing timestamp at the moment of this puzzle creation.
	UnixNanos time.Duration `msgpack:"t0"`
	// ServerHash if the first hash that let the server validate tour finish.
	ServerHash [32]byte `msgpack:"h0"`
	// IntegrityHash is a HMAC hash to let next gtp perform validation.
	IntegrityHash [32]byte `msgpack:"m0"`
}

// AnotherGuidePuzzle is another gtp's puzzle info given to the client on the
// current stop.
type AnotherGuidePuzzle struct {
	// GuideHash is the first hash that let the server validate tour finish.
	GuideHash [32]byte `msgpack:"h_s"`
	// IntegrityHash is a HMAC hash to let next gtp perform validation.
	IntegrityHash [32]byte `msgpack:"m_s"`
	// NextGuideAddr is the ID of the next tour gtp to follow.
	NextGuideID int `msgpack:"i_{s+1}"`
	// NextGuideAddr is the IP address of the next tour gtp to follow.
	NextGuideAddr net.IP `msgpack:"ip_{s+1}"`
	// UnixNanos coursing timestamp at the moment of this puzzle creation.
	UnixNanos time.Duration `msgpack:"t_s"`
}

// PuzzleSolvingMsg is another puzzle which client sends to the next gtp on
// the stop.
type PuzzleSolvingMsg struct {
	// The first tour hash to validate successful tour finish by a client.
	ServerHash [32]byte `msgpack:"h0"`
	// TourLength is decided by the main server to adjust the puzzle difficulty.
	// How many guided tours (count of round trips) will crypto have.
	TourLength uint16 `msgpack:"L"`
	// StopNumber is the passed stop (gtp) number by a client.
	StopNumber byte `msgpack:"s"`
	// PreviousGuideTS is the previous gtp puzzle generating timestamp.
	PreviousGuideTS time.Duration `msgpack:"t_{s-1}"`
	// PreviousIntegrityHash is the previous integrity HMAC hash from the previous gtp.
	PreviousIntegrityHash [32]byte `msgpack:"m_{s-1}"`
	// PassedGuides is a sequence of the IDs which client has been passed.
	PassedGuides []int `msgpack:"i_1,..,i_s"`
}
