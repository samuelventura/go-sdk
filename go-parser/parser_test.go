package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go test

func TestParser(t *testing.T) {
	root := parse("x^2+1")
	assert.Equal(t, "+", root.expr)
	assert.Equal(t, "^", root.left.expr)
	assert.Equal(t, "x", root.left.left.expr)
	assert.Equal(t, "2", root.left.right.expr)
	assert.Equal(t, "+1", root.right.expr)
}

func TestEvaluator(t *testing.T) {
	assert.Equal(t, 25.0, evald("x^2", 5))
	assert.Equal(t, 36.0, evald("x^2", 6))
	assert.Equal(t, 8.0, evald("x^3", 2))
	ast := parse("+x^2+2*x+1")
	fmt.Println(evals(ast))
	assert.Equal(t, 9.0, eval(ast, 2))
	assert.Equal(t, 16.0, eval(ast, 3))
	ast2 := parse("+x^4-x^2-1")
	fmt.Println(evals(ast2))
	assert.Equal(t, 11.0, eval(ast2, 2))
}
