package task

import (
	"crypto/rand"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IDGeneratorTestSuite struct {
	suite.Suite
	generator *TimeBasedIDGenerator
}

func (s *IDGeneratorTestSuite) SetupTest() {
	s.generator = &TimeBasedIDGenerator{}
}

func TestTimeBasedIDGenerator_GenerateID(t *testing.T) {
	generator := &TimeBasedIDGenerator{}
	ids := make(map[uint]bool)
	for i := 0; i < 1000; i++ {
		id := generator.GenerateID()
		assert.NotEqual(t, uint(0), id, "Generated ID should not be zero")
		assert.False(t, ids[id], "Generated ID %d is not unique", id)
		ids[id] = true

		time.Sleep(time.Microsecond)
	}
}

func TestTimeBasedIDGenerator_ConcurrentGeneration(t *testing.T) {
	generator := &TimeBasedIDGenerator{}
	numGoroutines := 100
	idsPerGoroutine := 100
	idsChan := make(chan uint, numGoroutines*idsPerGoroutine)

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < idsPerGoroutine; j++ {
				id := generator.GenerateID()
				idsChan <- id
			}
		}()
	}
	wg.Wait()
	close(idsChan)

	ids := make(map[uint]bool)
	for id := range idsChan {
		assert.False(t, ids[id], "Generated ID %d is not unique", id)
		ids[id] = true
	}
}

func TestIDGeneratorSuite(t *testing.T) {
	suite.Run(t, new(IDGeneratorTestSuite))
}

func (s *IDGeneratorTestSuite) TestGenerateIDWithRandomError() {
	originalRandRead := rand.Reader
	defer func() { rand.Reader = originalRandRead }()

	mockReader := &mockErrorReader{err: errors.New("random error")}
	rand.Reader = mockReader

	id := s.generator.GenerateID()
	assert.NotEqual(s.T(), uint(0), id)
}

type mockErrorReader struct {
	err error
}

func (m *mockErrorReader) Read(p []byte) (n int, err error) {
	return 0, m.err
}
