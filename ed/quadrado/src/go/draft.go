package main

import "fmt"

func calc_quadrado(x int) int {

	// x^y = (x-1)^y + 2*(x-1) + 1

	if x == 1 {
		fmt.Println("1^2 = 1")
		return 1
	}

	anterior := x - 1
	fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", x, anterior, anterior)
	quadrado := calc_quadrado(anterior) + 2*(anterior) + 1

	fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", x, anterior, anterior, quadrado)

	return quadrado
}

func main() {

	var x int

	fmt.Scan(&x)

	calc_quadrado(x)

}
