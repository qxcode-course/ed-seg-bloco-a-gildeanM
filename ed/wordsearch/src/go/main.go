package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x int
	y int
}

func Acima(i int, j int) Pos {
	return Pos{x: i - 1, y: j}
}

func Abaixo(i int, j int) Pos {
	return Pos{x: i + 1, y: j}
}

func Aesquerda(i int, j int) Pos {
	return Pos{x: i, y: j - 1}
}

func Adireita(i int, j int) Pos {
	return Pos{x: i, y: j + 1}
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	_, _ = grid, word

	wordLetters := []byte(word)
	idxLetterTest := 0

	neighsVisited := make(map[Pos]bool)

	for j := range grid[0] {
		neighsVisited[Pos{x: -1, y: j}] = true
	}

	for j := range grid[0] {
		neighsVisited[Pos{x: len(grid), y: j}] = true
	}

	for i := range grid {
		neighsVisited[Pos{x: i, y: -1}] = true
	}

	for i := range grid {
		neighsVisited[Pos{x: i, y: len(grid[0])}] = true
	}

	stack := NewStack[Pos]()

	for i, linha := range grid {
		for j, valor := range linha {

			atualPos := Pos{x: i, y: j}

			if wordLetters[idxLetterTest] == valor {
				stack.Push(atualPos)
				neighsVisited[atualPos] = true
			}

			var newPos Pos

			if !neighsVisited[Acima(i, j)] {
				newPos = Acima(i, j)
				i = newPos.x
				j = newPos.y
			} else if !neighsVisited[Adireita(i, j)] {
				newPos = Adireita(i, j)
				i = newPos.x
				j = newPos.y
			} else if !neighsVisited[Abaixo(i, j)] {
				newPos = Abaixo(i, j)
				i = newPos.x
				j = newPos.y
			} else if !neighsVisited[Aesquerda(i, j)] {
				newPos = Aesquerda(i, j)
				i = newPos.x
				j = newPos.y
			} else {
				return false
			}

		}

	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
