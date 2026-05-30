package main

import "fmt"

func balanced(str string) bool {

	stack := NewStack[rune]()

	for _, atualLetter := range str {

		if stack.IsEmpty() {
			stack.Push(atualLetter)
			continue
		}

		letterAtTop := stack.Top()

		if equals(letterAtTop, atualLetter) {
			stack.Pop()
		} else {
			stack.Push(atualLetter)
		}
	}

	if !stack.IsEmpty() {
		return false
	}

	return true
}

func equals(l1 rune, l2 rune) bool {

	if l1 == '(' && l2 == ')' {
		return true
	}

	if l1 == '[' && l2 == ']' {
		return true
	}

	return false
}

func main() {

	var str string

	fmt.Scan(&str)

	if balanced(str) {
		fmt.Println("balanceado")
	} else {
		fmt.Println("nao balanceado")
	}

}
