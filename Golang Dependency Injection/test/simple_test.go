package test

import (
	"fmt"
	"golang_dependency_injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService(true)

	fmt.Println(err)
	fmt.Println(simpleService)

	assert.Nil(t, simpleService)
}
