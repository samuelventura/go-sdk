package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//go test
func TestParser(t *testing.T) {
	assert.Equal(t, 25.0, Evaluate("x^2", 5))
	assert.Equal(t, 36.0, Evaluate("x^2", 6))
	assert.Equal(t, 8.0, Evaluate("x^3", 2))
}
