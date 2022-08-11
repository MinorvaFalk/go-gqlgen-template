package ulid

import (
	"crypto/rand"

	"github.com/oklog/ulid/v2"
)

type ID string

var defaultEntropySource *ulid.MonotonicEntropy

func init() {
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}
