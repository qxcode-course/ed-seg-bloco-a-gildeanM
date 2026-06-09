package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Search node with value
// Starts with node == root
// Value is the value we are looking for
// Returns the node with the value or nil if not found
func find(node *Node, value int) *Node {

	var fnode *Node

	if node == nil {
		return node
	}

	if node.Value == value {
		return node
	}

	fnode = find(node.Left, value)

	if fnode != nil {
		return fnode
	}
	fnode = find(node.Right, value)

	return fnode
}

// node is the node we want to find the height of
// the height of a node is the number of edges on the longest path from the node to a leaf
func getHeight(node *Node) int {

	leftheight := 0
	rightheight := 0

	if node == nil {
		return 0
	}

	leftheight += 1
	leftheight += getHeight(node.Left)
	rightheight += 1
	rightheight += getHeight(node.Right)

	if leftheight > rightheight {
		return leftheight
	} else {
		return rightheight
	}

}

// node is the root of the tree
// level is the current level in the tree (1 for root)
// value is the value we are looking for
func calcNodeDepth(node *Node, level int, value int) int {

	leftdepth := level
	rightdepth := level

	if node == nil {
		return 1
	}

	if node.Value == value {
		return level
	}

	leftdepth = calcNodeDepth(node.Left, leftdepth+1, value)
	rightdepth = calcNodeDepth(node.Right, rightdepth+1, value)

	if rightdepth > leftdepth {
		return rightdepth
	} else {
		return leftdepth
	}

}

// --------------------------------------------------------------------
// Don't change from here
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

	line := scanner.Text()
	parts := strings.Split(line, " ")
	root := create(&parts)

	scanner.Scan()
	line = scanner.Text()
	fmt.Println("Arvore:")
	BShow(root, "")

	values := strings.FieldsSeq(line)
	for s := range values {
		value, _ := strconv.Atoi(s)
		node := find(root, value)
		if node != nil {
			fmt.Printf("Altura: %d, Profundidade: %d\n", getHeight(node), calcNodeDepth(root, 1, value))
		} else {
			fmt.Println("-1")
		}
	}
}
