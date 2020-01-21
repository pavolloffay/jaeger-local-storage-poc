package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewField(t *testing.T) {
	field := NewField("field1", StringType, []byte("value1"))
	assert.Equal(t, "field1", field.Name())
	assert.Equal(t, StringType, field.Type())
	assert.Equal(t, "value1", string(field.Value()))
}
