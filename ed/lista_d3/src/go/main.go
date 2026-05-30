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
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func equals(list1 *LList, list2 *LList) bool {

	n1 := list1.Front()
	n2 := list2.Front()
	for {

		if n1 == nil && n2 == nil {
			break
		}

		if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
			return false
		}

		if n1.Value != n2.Value {
			return false
		}

		n1 = n1.Next()
		n2 = n2.Next()
	}

	return true
}

func (n *Node) Next() *Node {

	if n.next == n.root {
		return nil
	}

	return n.next
}

func (ll LList) Back() *Node {

	if ll.Size() == 0 {
		return nil
	}

	return ll.root.prev
}

func (ll LList) Front() *Node {

	if ll.Size() == 0 {
		return nil
	}

	return ll.root.next
}

func (n *Node) Prev() *Node {

	if n.prev == n.root {
		return nil
	}

	return n.prev
}

func (ll LList) Size() int {
	return ll.size
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark

	mark.prev.next = n
	mark.prev = n

	l.size++
}

func (ll *LList) PushFront(value int) {
	ll.insertAfter(ll.root, value)
}

func (ll *LList) insertAfter(mark *Node, value int) {
	newNode := &Node{
		Value: value,
		next:  ll.root.next,
		prev:  ll.root,
		root:  ll.root,
	}

	ll.root.next.prev = newNode
	ll.root.next = newNode

	ll.size++
}

func addsorted(ll *LList, value int) {

	var insert *Node
	for node := ll.Front(); node != nil; node = node.Next() {
		if node.Value > value {
			insert = node
			break
		}
	}

	if insert != nil {
		ll.Insert(insert, value)
	} else {
		ll.PushBack(value)
	}

}

func (ll *LList) Insert(node *Node, value int) {
	newNode := &Node{
		Value: value,
		next:  node,
		prev:  node.prev,
		root:  ll.root,
	}

	node.prev.next = newNode
	node.prev = newNode

	ll.size++

}

func reverse(ll *LList) {
	for node := ll.Front(); node != nil; {
		next := node.Next()
		node.next, node.prev = node.prev, node.next
		node = next
	}

	ll.root.next, ll.root.prev = ll.root.prev, ll.root.next
}

func merge(lla *LList, llb *LList) *LList {

	llmerged := NewLList()

	for node := lla.Front(); node != nil; node = node.Next() {
		addsorted(llmerged, node.Value)
	}

	for node := llb.Front(); node != nil; node = node.Next() {
		addsorted(llmerged, node.Value)
	}

	return llmerged
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}

func (ll LList) String() string {
	out := ""

	out += "["
	for node := ll.Front(); node != nil; node = node.Next() {
		out += strconv.Itoa(node.Value)
		if node.Next() != nil {
			out += ", "
		}
	}
	out += "]"

	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
