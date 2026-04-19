package Golang_Generic

import (
	"fmt"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (data *MyData[T]) GetValue() T {
	return data.Value
}

func (data *MyData[T]) SetValue(value T) {
	data.Value = value
}

func TestGetterSetter(t *testing.T) {
	data := MyData[string]{}
	data.SetValue("Ahmad")
	fmt.Println(data.GetValue())
	fmt.Println(data)

	ChangeValue(&data, "Gayuh")
	fmt.Println(data)
}
