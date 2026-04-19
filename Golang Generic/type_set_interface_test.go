package Golang_Generic

import (
	"fmt"
	"testing"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	fmt.Println(Min[int](10, 11))
	fmt.Println(Min[int](100, 1100))
	fmt.Println(Min(100, 1100))
	fmt.Println(Min[Age](100, 1100))
}
