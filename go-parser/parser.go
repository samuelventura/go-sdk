package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type node struct {
	expr  string
	left  *node
	right *node
}

func trace(args ...interface{}) {
	fmt.Println(args...)
}

func parse(expr string) *node {
	trace("parse", expr)
	pc := 0
	expr = strings.TrimSpace(expr)
	for i, r := range expr {
		if r == '(' {
			pc++
		}
		if r == ')' {
			pc--
		}
		if pc > 0 {
			continue
		}
		if unicode.IsSpace(r) {
			continue
		}
		//ignore first sign
		if i == 0 {
			switch r {
			case '+':
				continue
			case '-':
				continue
			}
		}
		switch r {
		case '+':
			n := &node{}
			n.expr = "+"
			n.left = parse_m(expr[:i])
			n.right = parse(expr[i:])
			return n
		case '-':
			n := &node{}
			n.expr = "+"
			n.left = parse_m(expr[:i])
			n.right = parse(expr[i:])
			return n
		}
	}
	return parse_m(expr)
}

func parse_m(expr string) *node {
	trace("parse_m", expr)
	s := 0
	pc := 0
	expr = strings.TrimSpace(expr)
	for i, r := range expr {
		if r == '(' {
			pc++
		}
		if r == ')' {
			pc--
		}
		if pc > 0 {
			continue
		}
		if unicode.IsSpace(r) {
			continue
		}
		if i == 0 {
			switch r {
			//remove + from front of x
			case '+':
				s = 1
				continue
			case '-':
				n := &node{}
				n.expr = "*"
				n.left = &node{expr: "-1"}
				n.right = parse_m(expr[i+1:])
				return n
			}
		}
		switch r {
		case '*':
			n := &node{}
			n.expr = string(r)
			n.left = parse_d(expr[s:i])
			n.right = parse_d(expr[i+1:])
			return n
		}
	}
	return parse_d(expr)
}

func parse_d(expr string) *node {
	trace("parse_d", expr)
	pc := 0
	expr = strings.TrimSpace(expr)
	for i, r := range expr {
		if r == '(' {
			pc++
		}
		if r == ')' {
			pc--
		}
		if pc > 0 {
			continue
		}
		if unicode.IsSpace(r) {
			continue
		}
		switch r {
		case '/':
			n := &node{}
			n.expr = string(r)
			n.left = parse_p(expr[0:i])
			n.right = parse_p(expr[i+1:])
			return n
		}
	}
	return parse_p(expr)
}

func parse_p(expr string) *node {
	trace("parse_p", expr)
	pc := 0
	expr = strings.TrimSpace(expr)
	for i, r := range expr {
		if r == '(' {
			pc++
		}
		if r == ')' {
			pc--
		}
		if pc > 0 {
			continue
		}
		if unicode.IsSpace(r) {
			continue
		}
		switch r {
		case '^':
			n := &node{}
			n.expr = string(r)
			n.left = parse_l(expr[0:i])
			n.right = parse_l(expr[i+1:])
			return n
		}
	}
	return parse_l(expr)
}

func parse_l(expr string) *node {
	trace("parse_l", expr)
	l := len(expr)
	if expr[0] == '(' && expr[l-1] == ')' {
		return parse(expr[1 : l-1])
	}
	return &node{expr: expr}
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
		case "/":
			return left / right
		case "^":
			return math.Pow(left, right)
		default:
			panic("Unknown operation")
		}
	}
	expr := trim(n.expr)
	switch expr {
	case "x":
		return value
	case "+x": //needed
		return value
	default:
		val, err := strconv.ParseFloat(expr, 64)
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

func trim(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
