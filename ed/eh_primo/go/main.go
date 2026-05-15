package main

import (
	"fmt"
	"math"
)

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	_, _ = x, div

	if x <= 1 {
		return false
	}
	if x <= 3 {
		return true
	}
	if x%2 == 0 || x%3 == 0 {
		return false
	}

	if div > int(math.Sqrt(float64(x)))+1 {
		return true
	} else {
		if x%div == 0 || x%(div+2) == 0 {
			return false
		}
		return eh_primo(x, div+6)
	}

}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 5) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
