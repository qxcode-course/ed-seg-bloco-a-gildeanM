package main

import "fmt"

// func padrao_f3(n int) int {

// 	if n == 1 {
// 		return 1
// 	}

// 	return padrao_f3(n-1) + n

// }

// func padrao_f4(n int) int {

// 	if n == 1 {
// 		return 1
// 	}

// 	return padrao_f4(n-1) + (2*n - 1)

// }

// func padrao_f5(n int) int {

// 	if n == 1 {
// 		return 1
// 	}

// 	return padrao_f5(n-1) + n*3 - 2

// }

// func padrao_f6(n int) int {

// 	if n == 1 {
// 		return 1
// 	}

// 	return padrao_f6(n-1) + (n*4 - 3)

// 	// grau 2 -> n*2 + 1 (n - 1) -> n*3 - 1 -> 1 + 5 (n*2 + (n - 1) + (n - 2))
// 	// grau 3 -> n*2 + 3 (n + 0) -> n*3 + 0 -> (1 + 5) + 9 (n*2 + (n-1) + (n - 2))
// 	// grau 4 -> n*2 + 5 (n + 1) -> n*3 + 1 -> ((1 + 5) + 9) + 13 (n*2 + (n-1) + (n - 2))
// 	// grau 5 -> n*2 + 7 (n + 2) -> n*3 + 2 -> (((1 + 5) + 9) + 13) + 17 (n*2 + (n-1) + (n - 2))

// }

// func padrao_f7(n int) int {

// 	if n == 1 {
// 		return 1
// 	}

// 	return padrao_f7(n-1) + (n*2 + (n - 1) + (n - 2))

// 	// grau 2 -> 1 + 6
// 	// grau 3 -> (1 + 6) +
// 	// grau 4 ->
// 	// grau 5 ->

// }

func padroes(seq int, n int) int {

	if n == 1 {
		return 1
	}

	return padroes(seq, n-1) + (seq-2)*n - (seq - 3)

}

func main() {

	var seq, n, pontos int

	fmt.Scan(&seq, &n)

	pontos = padroes(seq, n)

	// switch seq {
	// case 3:
	// 	pontos = padrao_f3(n)
	// case 4:
	// 	pontos = padrao_f4(n)
	// case 5:
	// 	pontos = padrao_f5(n)
	// case 6:
	// 	pontos = padrao_f6(n)
	// }

	fmt.Println(pontos)
}
