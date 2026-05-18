package main

import "fmt"

// func dobra_doacao(entradas float32, doacao float32) float32 {

// 	// fmt.Printf("%f - %f\n", entradas, doacao)

// 	if entradas == 1 {
// 		return doacao / 2
// 	}

// 	return (dobra_doacao(entradas-1, doacao) + doacao) / 2

// 	// entra -> 17,50
// 	// dobra (17,50 * 2) - 20
// 	// ((17,50 * 2) - 20)*2 - 20
// 	// (((17,50 * 2) - 20)*2 - 20)*2 - 20 = 0
// 	// (0 + 20)/2 = 10
// 	// (10 + 20)/2 = 15
// 	// (15 + 20)/2 = 17,5

// }

func main() {

	var entradas, doacao float32

	var dobra_doacao func(entradas float32) float32

	dobra_doacao = func(entradas float32) float32 {

		if entradas == 1 {
			return doacao / 2
		}

		return (dobra_doacao(entradas-1) + doacao) / 2

	}

	fmt.Scan(&entradas, &doacao)

	fmt.Printf("%.2f\n", dobra_doacao(entradas))
}
