package Golang_Generic

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	fmt.Println(param1)
	fmt.Println(param2)
	return param1, param2
}

func TestMultipleParameter(t *testing.T) {
	t1, t2 := MultipleParameter[int, string](123, "Test")
	fmt.Println(t1, t2)
}
