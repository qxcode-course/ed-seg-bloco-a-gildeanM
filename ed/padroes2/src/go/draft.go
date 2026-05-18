package main

import "fmt"

// grau 1 => 1 - 1 - 1
// grau 2 => 1 + (1) - 1 + (2 + 1) - 2 + (1)
// grau 3 => 2 + (1) - 4 + (2 + 3)  - 2 + (1)

func padroes(n int) int {

	if n == 1 {
		return n * 3
	}

	return padroes(n-1) + (n + n - 1) + 2

}

func main() {

	var n int

	fmt.Scan(&n)

	fmt.Println(padroes(n))
}
