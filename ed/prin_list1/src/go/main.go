package main

import (
	"fmt"
	"strconv"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	out := ""

	out += "[ "
	for i := l.Front(); i != nil; i = i.Next() {
		out += strconv.Itoa(i.Value)
		if i == sword {
			out += ">"
		}
		out += " "
	}
	out += "]"
	return out
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	next := it.Next()
	if next == it {
		next = l.root.next
	}
	return next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
