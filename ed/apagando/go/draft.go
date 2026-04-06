package main

import (
	"fmt"
)

func hasQuitted(quitters map[int]bool, person int) bool {
	if quitters[person] {
		return true
	}
	return false
}

func main() {

	var qtd_persons int
	fmt.Scanf("%d\n", &qtd_persons)

	queue := make([]int, qtd_persons)

	for i := 0; i < qtd_persons; i++ {
		fmt.Scan(&queue[i])
	}

	var qtd_quitters int
	fmt.Scan(&qtd_quitters)

	quitters := make(map[int]bool, qtd_quitters)

	for i := 0; i < qtd_quitters; i++ {
		var quitter int
		fmt.Scan(&quitter)
		quitters[quitter] = true
	}

	qtdPersonsAfterQuitters := qtd_persons - qtd_quitters
	queueAfterQuitters := make([]int, 0, qtdPersonsAfterQuitters)

	for _, person := range queue {
		if !quitters[person] {
			queueAfterQuitters = append(queueAfterQuitters, person)
		}

	}

	for _, v := range queueAfterQuitters {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")

}
