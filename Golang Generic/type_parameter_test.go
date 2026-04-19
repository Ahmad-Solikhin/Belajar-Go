package Golang_Generic

import (
	"fmt"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	resultString := Length[string]("hello")
	fmt.Println(resultString)

	resultInt := Length[int](12345)
	fmt.Println(resultInt)
}
