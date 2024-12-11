package main

import (
	"fmt"
	"math"
)
func main() {
	var a, b  float64
	fmt.Scan(&a, &b)
	fmt.Println(Hypot(a,b))
}

func Hypot(a,b float64) float64{
	return math.Sqrt(a * a + b * b)
}