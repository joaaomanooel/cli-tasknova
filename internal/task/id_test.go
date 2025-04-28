package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeBasedIDGenerator_GenerateID(t *testing.T) {
	generator := &TimeBasedIDGenerator{}

	id1 := generator.GenerateID()
	id2 := generator.GenerateID()

	assert.NotEqual(t, id1, id2, "Generated IDs should be unique")
	assert.Greater(t, id1, uint(0), "Generated ID should be greater than 0")
	assert.Greater(t, id2, uint(0), "Generated ID should be greater than 0")
}
