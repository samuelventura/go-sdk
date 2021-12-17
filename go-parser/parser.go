package main

import (
	"math"
	"strconv"
	"strings"
)

type node struct {
	expr  string
	left  *node
	right *node
}

func parse(expr string) *node {
	n := &node{}
	for i, c := range expr {
		if i == 0 {
			continue
		}
		switch c {
		case '+':
			n.expr = "+"
			n.left = parsei(expr[:i])
			n.right = parse(expr[i:])
			return n
		case '-':
			n.expr = "+"
			n.left = parsei(expr[:i])
			n.right = parse(expr[i:])
			return n
		}
	}
	return parsei(expr)
}

func parsei(expr string) *node {
	n := &node{}
	s := 0
	for i, c := range expr {
		if i == 0 {
			switch c {
			case '+':
				s = 1
				continue
			case '-':
				n.expr = "*"
				n.left = &node{expr: "-1"}
				n.right = parsei(expr[i+1:])
				return n
			}
		}
		switch c {
		case '*':
			n.expr = string(c)
			n.left = parsei(expr[s:i])
			n.right = parsei(expr[i+1:])
			return n
		case '^':
			n.expr = string(c)
			n.left = parsei(expr[s:i])
			n.right = parsei(expr[i+1:])
			return n
		}
	}
	n.expr = expr
	return n
}

func evald(expr string, value float64) float64 {
	root := parse(expr)
	return eval(root, value)
}

func eval(n *node, value float64) float64 {
	if n.left != nil && n.right != nil {
		left := eval(n.left, value)
		right := eval(n.right, value)
		switch n.expr {
		case "+":
			return left + right
		case "*":
			return left * right
		case "^":
			return math.Pow(left, right)
		default:
			panic("Unknown operation")
		}
	}
	switch n.expr {
	case "x":
		return value
	case "+x":
		return value
	default:
		val, err := strconv.ParseFloat(n.expr, 64)
		if err != nil {
			panic(err)
		}
		return val
	}
}

func evals(n *node) string {
	var b strings.Builder
	if n.left != nil && n.right != nil {
		b.WriteString("(")
		b.WriteString(n.expr)
		b.WriteString(" ")
		b.WriteString(evals(n.left))
		b.WriteString(" ")
		b.WriteString(evals(n.right))
		b.WriteString(")")
		return b.String()
	}
	b.WriteString(n.expr)
	return b.String()
}
