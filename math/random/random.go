// Package random provides .NET System.Random-like random number generation.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.random?view=netframework-4.7.2
package random

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"time"
)

// Random represents a pseudo-random number generator.
// Equivalent to System.Random in .NET.
type Random struct {
	rng *rand.Rand
}

// NewRandom creates a new Random with a time-based seed.
func NewRandom() *Random {
	return &Random{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// NewRandomWithSeed creates a new Random with the specified seed.
func NewRandomWithSeed(seed int64) *Random {
	return &Random{
		rng: rand.New(rand.NewSource(seed)),
	}
}

// Next returns a non-negative random integer.
func (r *Random) Next() int {
	return r.rng.Int()
}

// NextN returns a random integer in [0, n).
func (r *Random) NextN(n int) int {
	if n <= 0 {
		return 0
	}
	return r.rng.Intn(n)
}

// NextRange returns a random integer in [min, max).
func (r *Random) NextRange(min, max int) int {
	if min >= max {
		return min
	}
	return min + r.rng.Intn(max-min)
}

// NextDouble returns a random float64 in [0.0, 1.0).
func (r *Random) NextDouble() float64 {
	return r.rng.Float64()
}

// NextBytes fills the provided byte slice with random bytes.
func (r *Random) NextBytes(buffer []byte) {
	r.rng.Read(buffer)
}

// ---- Cryptographic random (static methods) ----

// CryptoNextInt returns a cryptographically secure random integer.
func CryptoNextInt() (int64, error) {
	n, err := crand.Int(crand.Reader, big.NewInt(1<<62))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

// CryptoNextBytes fills buffer with cryptographically secure random bytes.
func CryptoNextBytes(buffer []byte) error {
	_, err := crand.Read(buffer)
	return err
}
