package task

import (
	"encoding/binary"
	"time"
)

type IDGenerator interface {
	GenerateID() uint
}

type TimeBasedIDGenerator struct{}

// func (g *TimeBasedIDGenerator) GenerateID() uint {
// 	now := time.Now()

// 	seconds := int32(now.Unix())
// 	nanos := int32(now.Nanosecond())

// 	randBytes := make([]byte, 4)
// 	if _, err := rand.Read(randBytes); err != nil {
// 		timeComponent := (uint(seconds&0xFFFF) | uint((nanos&0xFFFF)<<16))
// 		return timeComponent
// 	}

// 	random := binary.BigEndian.Uint32(randBytes)

// 	timeComponent := uint(seconds&0xFFFF) | uint((nanos&0xFFFF)<<16)
// 	return timeComponent ^ uint(random)
// }

func (g *TimeBasedIDGenerator) GenerateID() uint {
	timestamp := uint(time.Now().UnixNano())

	randBytes := make([]byte, 4)
	random := binary.BigEndian.Uint32(randBytes)

	return timestamp ^ uint(random)
}
