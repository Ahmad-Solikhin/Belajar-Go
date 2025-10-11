package helper

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Before run")

	m.Run()

	fmt.Println("After run")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Gayuh")

	if result != "Hello Gayuh" {
		panic("Result not same")
	}
}

func TestHelloWorldAhmad(t *testing.T) {
	result := HelloWorld("Ahmad")

	if result != "Hello Ahmad" {
		panic("Result not same")
	}
}
