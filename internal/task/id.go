package task

import (
    "crypto/rand"
    "encoding/binary"
    "time"
)

type IDGenerator interface {
    GenerateID() uint
}

type TimeBasedIDGenerator struct{}

func (g *TimeBasedIDGenerator) GenerateID() uint {
    // Combine timestamp with random bytes for uniqueness
    timestamp := uint(time.Now().UnixNano())
    
    // Generate 4 random bytes
    randBytes := make([]byte, 4)
    rand.Read(randBytes)
    random := binary.BigEndian.Uint32(randBytes)
    
    // Combine timestamp and random number
    return timestamp ^ uint(random)
}