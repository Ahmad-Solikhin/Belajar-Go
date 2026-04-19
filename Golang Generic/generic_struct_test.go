package Golang_Generic

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func (t *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (t *Data[T]) ChangeFirst(first T) T {
	t.First = first
	return t.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "hello",
		Second: "world",
	}

	fmt.Println(data.SayHello("Gayuh"))

	fmt.Println(data.ChangeFirst("Ahmad"))

	fmt.Println(data)
}
