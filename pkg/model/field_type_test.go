package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvBoolToBytes(t *testing.T) {
	restrue := ConvBoolToBytes(true)
	resfalse := ConvBoolToBytes(false)
	assert.Len(t, restrue, 1)
	assert.Len(t, resfalse, 1)
	assert.Equal(t, byte(1), restrue[0])
	assert.Equal(t, byte(0), resfalse[0])
}
