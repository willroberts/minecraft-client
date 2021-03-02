package minecraft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	gen := &idGenerator{}

	assert.Equal(t, gen.GenerateID(), int32(1))
	assert.Equal(t, gen.GenerateID(), int32(2))
	assert.Equal(t, gen.GenerateID(), int32(3))
}
