package Golang_Generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int16 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
}

func GetFirst[T []E, E any](slice T) E {
	return slice[0]
}

func TestGetFirst(t *testing.T) {
	list := []int{1, 2, 3}
	fmt.Println(GetFirst(list))
}
