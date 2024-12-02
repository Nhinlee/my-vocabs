package genid

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid"
)

func NewULID() string {
	// Seed the ULID generator with the current time in milliseconds
	entropy := ulid.Monotonic(rand.Reader, 0)
	ulid, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
	if err != nil {
		// Handle the error, return an empty string or log it
		panic(err) // Or handle it according to your use case
	}
	return ulid.String()
}
