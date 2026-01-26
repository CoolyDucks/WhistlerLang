package main

import (
	"fmt"
	"math"
	"strconv"
)

func Add(a, b string) float64 {
	x, _ := strconv.ParseFloat(a, 64)
	y, _ := strconv.ParseFloat(b, 64)
	return x + y
}

func Sub(a, b string) float64 {
	x, _ := strconv.ParseFloat(a, 64)
	y, _ := strconv.ParseFloat(b, 64)
	return x - y
}

func Mul(a, b string) float64 {
	x, _ := strconv.ParseFloat(a, 64)
	y, _ := strconv.ParseFloat(b, 64)
	return x * y
}

func Div(a, b string) float64 {
	x, _ := strconv.ParseFloat(a, 64)
	y, _ := strconv.ParseFloat(b, 64)
	if y == 0 {
		fmt.Println("Division by zero")
		return 0
	}
	return x / y
}

func Pow(a, b string) float64 {
	x, _ := strconv.ParseFloat(a, 64)
	y, _ := strconv.ParseFloat(b, 64)
	return math.Pow(x, y)
}
