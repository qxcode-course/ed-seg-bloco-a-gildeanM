package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x, y int
}

func atualizarCobra(cobra []Pos, direcao byte) {
	for i := len(cobra) - 1; i > 0; i-- {
		cobra[i] = cobra[i-1]
	}

	switch direcao {
	case 'U':
		cobra[0].y -= 1
	case 'D':
		cobra[0].y += 1
	case 'L':
		cobra[0].x -= 1
	case 'R':
		cobra[0].x += 1
	}
}

func main() {
	var qtdGomos int
	var pos byte

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %c\n", &qtdGomos, &pos)

	cobra := make([]Pos, qtdGomos)

	for i := 0; i < qtdGomos; i++ {
		fmt.Fscanf(reader, "%d %d\n", &cobra[i].x, &cobra[i].y)
	}

	atualizarCobra(cobra, pos)

	for i := 0; i < qtdGomos; i++ {
		fmt.Printf("%d %d\n", cobra[i].x, cobra[i].y)
	}
}
