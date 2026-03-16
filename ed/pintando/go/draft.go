package main

import (
	"fmt"
	"math"
)

func main() {
	// puts("qxcode");
	var a, b, c float64
	var p, area float64

	fmt.Scanf("%f\n%f\n%f", &a, &b, &c)

	p = (a + b + c) / 2

	area = math.Sqrt(
		p *
			(p - a) *
			(p - b) *
			(p - c),
	)

	fmt.Printf("%.2f\n", area)
}
