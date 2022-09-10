package crypto

import (
	"crypto/rand"
	"math"

	"github.com/pkg/errors"
)

// GenerateRandomUint16 generates random unsigned 16-bit integer from uniform
// distribution.
func GenerateRandomUint16() (uint16, error) {
	buf := []byte{0, 0}
	if _, err := rand.Reader.Read(buf); err != nil {
		return 0, errors.Wrap(err, "can't generate random bytes")
	}

	return uint16(buf[0]) % math.MaxUint16, nil
}
