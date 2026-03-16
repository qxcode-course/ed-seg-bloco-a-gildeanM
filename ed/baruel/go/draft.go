package main

import (
	"bufio"
	"fmt"
	"os"
)

func fill(arr []int, size int, value int) {
	for i := 0; i < size; i++ {
		arr[i] = value
	}
}

func add_fig_missing(arr []int, next_index *int, fig int) {
	arr[*next_index] = fig
	(*next_index)++
}

func add_figs_repeated(arr []int, next_index *int, fig int, qtd_repeat int) {
	for i := 0; i < qtd_repeat; i++ {
		arr[*next_index] = fig
		(*next_index)++
	}
}

func print_fig_missing(arr []int, size int) {
	if size == 1 {
		fmt.Println("N")
		return
	}

	for i := 1; i < size; i++ {
		fmt.Print(arr[i])
		if i == size-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}

func print_fig_repeated(arr []int, size int) {
	if size == 0 {
		fmt.Println("N")
		return
	}

	for i := 0; i < size; i++ {
		fmt.Print(arr[i])
		if i == size-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}

func main() {
	// puts("qxcode");

	in := bufio.NewReader(os.Stdin)
	var qtdAlbum, qtdBaruelHas int
	fmt.Fscan(in, &qtdAlbum, &qtdBaruelHas)

	qtdCanRepeat := qtdAlbum * 2

	figRepeated := make([]int, qtdCanRepeat)
	figMissing := make([]int, qtdAlbum)

	figurinhas := make([]int, qtdAlbum+1)

	fill(figurinhas, qtdAlbum+1, 0)

	for i := 1; i <= qtdBaruelHas; i++ {
		var fig int
		fmt.Fscan(in, &fig)
		figurinhas[fig]++
	}

	next_index_missing := 0
	next_index_repeated := 0
	for i := 0; i <= qtdAlbum; i++ {
		if figurinhas[i] == 0 {
			add_fig_missing(figMissing, &next_index_missing, i)
		} else if figurinhas[i] > 1 {
			add_figs_repeated(figRepeated, &next_index_repeated, i, figurinhas[i]-1)
		}
	}

	print_fig_repeated(figRepeated, next_index_repeated)
	print_fig_missing(figMissing, next_index_missing)
}
