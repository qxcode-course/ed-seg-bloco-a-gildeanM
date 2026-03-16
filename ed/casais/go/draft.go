package main

import (
	"bufio"
	"fmt"
	"os"
)

func match_pair(animal int, descasados []int, n int) int {
	for i := 0; i < n; i++ {
		pair := 0 - animal
		if descasados[i] == pair {
			return i
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// puts("qxcode");

	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	descasados := make([]int, n)
	qtd_pairs := 0

	for i := 0; i < n; i++ {
		var animal int
		fmt.Fscan(reader, &animal)

		pair_index := match_pair(animal, descasados, i)

		if pair_index != -1 {
			descasados[pair_index] = 0
			qtd_pairs++
		} else {
			descasados[i] = animal
		}
	}

	fmt.Fprintln(writer, qtd_pairs)
}
