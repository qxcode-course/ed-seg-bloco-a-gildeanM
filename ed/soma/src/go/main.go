package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func rec_sum(node *Node) int {

	sum := 0
	if node == nil {
		return sum
	}

	sum += node.Value
	sum += rec_sum(node.Left)
	sum += rec_sum(node.Right)

	return sum
}

func rec_min(node *Node) int {
	min := math.MaxInt

	if node == nil {
		return min
	}

	if min > node.Value {
		min = node.Value
	}

	minleft := rec_min(node.Left)
	if min > minleft {
		min = minleft
	}

	minright := rec_min(node.Right)
	if min > minright {
		min = minright
	}

	return min
}

// MyShow imprime a árvore binária de forma formatada.
func MyShow(node *Node, nivel int) {
	_, _ = node, nivel
	// TODO
}

func BShow(node *Node, heranca string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, heranca+"l")
	}
	for i := 0; i < len(heranca)-1; i++ {
		if heranca[i] != heranca[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if heranca != "" {
		if heranca[len(heranca)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, heranca+"r")
	}
}

func create(parts *[]string) *Node {
	elem := (*parts)[0]
	*parts = (*parts)[1:]
	if elem == "#" {
		return nil
	}
	value, _ := strconv.Atoi(elem)
	node := &Node{Value: value}
	node.Left = create(parts)
	node.Right = create(parts)
	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	root := create(&parts)
	fmt.Println("Arvore:")
	BShow(root, "")
	fmt.Printf("Soma: %d, Minimo: %d\n", rec_sum(root), rec_min(root))
}
