package Golang_Generic

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBag(t *testing.T) {
	numbers := Bag[int]([]int{1, 2, 3})
	PrintBag(numbers)
}
