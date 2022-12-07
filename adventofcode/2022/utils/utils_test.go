package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	assert.Equal(t, 1, Atoi("1"))
	assert.Equal(t, 12345, Atoi("12345"))
}
