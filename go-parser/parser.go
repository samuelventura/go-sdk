package main

import (
	"math"
	"strconv"
	"strings"
)

func Evaluate(expr string, value float64) float64 {
	parts := strings.Split(expr, "^")
	expo, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		panic(err)
	}
	return math.Pow(value, expo)
}
