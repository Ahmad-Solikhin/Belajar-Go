package test

import (
	"golang_dependency_injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializeConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
