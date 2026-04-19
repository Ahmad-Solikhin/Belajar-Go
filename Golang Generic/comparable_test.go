package Golang_Generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	return value1 == value2
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Gayuh", "Gayuh"))
}
