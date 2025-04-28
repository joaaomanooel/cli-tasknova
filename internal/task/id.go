package task

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// IDGenerator defines the interface for generating unique task IDs
type IDGenerator interface {
	GenerateID() uint
}

// TimeBasedIDGenerator implements IDGenerator using time-based generation
type TimeBasedIDGenerator struct{}

func (g *TimeBasedIDGenerator) GenerateID() uint {
	// Get current timestamp with microsecond precision
	now := time.Now()
	timestamp := now.UnixMicro()

	// Generate 4 random bytes
	randBytes := make([]byte, 4)
	if _, err := rand.Read(randBytes); err != nil {
		// If random generation fails, use nanoseconds as fallback
		return uint(now.UnixNano() & 0xFFFFFFFF)
	}

	// Combine timestamp and random components
	random := binary.BigEndian.Uint32(randBytes)
	timeComponent := uint32(timestamp & 0xFFFFFFFF)

	// XOR the components to create final ID
	return uint(timeComponent ^ random)
}
