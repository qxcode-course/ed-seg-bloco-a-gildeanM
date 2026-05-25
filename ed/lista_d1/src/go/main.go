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
}

type LList struct {
	root *Node
}

func NewLList() *LList {

	root := &Node{}

	root.next = root
	root.prev = root

	return &LList{
		root: root,
	}
}

func (ll *LList) PushFront(value int) {
	node := &Node{Value: value}

	first := ll.root.next

	node.next = first
	node.prev = ll.root

	first.prev = node
	ll.root.next = node

}

func (ll *LList) PushBack(value int) {
	node := &Node{Value: value}

	last := ll.root.prev

	node.next = ll.root
	node.prev = last

	last.next = node
	ll.root.prev = node

}

func (ll *LList) PopFront() {

	if ll.root.next == ll.root {
		return
	}

	first := ll.root.next

	ll.root.next = first.next
	first.next.prev = ll.root

}

func (ll *LList) PopBack() {

	if ll.root.prev == ll.root {
		return
	}

	last := ll.root.prev

	last.prev.next = ll.root
	ll.root.prev = last.prev

}

func (ll *LList) Clear() {

	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func (ll LList) Size() int {

	count := 0

	node := ll.root.next
	for node != ll.root {
		count++
		node = node.next
	}

	return count

}

func (ll LList) String() string {

	output := "["

	node := ll.root.next
	for node != ll.root {

		output += strconv.Itoa(node.Value)

		if node.next != ll.root {
			output += ", "
		}

		node = node.next
	}
	output += "]"
	return output
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ll := NewLList()

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
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
