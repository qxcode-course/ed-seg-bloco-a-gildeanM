package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s Set) String() string {

	var vectorString strings.Builder
	vectorString.WriteByte('[')
	for i := 0; i < s.size; i++ {
		vectorString.WriteString(strconv.Itoa(s.data[i]))
		if i != s.size-1 {
			vectorString.WriteString(", ")
		}
	}
	vectorString.WriteByte(']')

	return vectorString.String()
}

func (s *Set) Reserve(newCapacity int) {

	if newCapacity <= s.capacity {
		return
	}

	newSet := NewSet(newCapacity)

	newSet.size = s.size

	for i := 0; i < s.size; i++ {
		newSet.data[i] = s.data[i]
	}

	*s = *newSet

}

func (s *Set) Insert(value int) {

	if s.Contains(value) {
		return
	}

	idx := s.size

	if value < s.data[0] {
		idx = 0
	} else {
		for i := 0; i < s.size-1; i++ {
			if value > s.data[i] && value < s.data[i+1] {
				idx = i + 1
				break
			}
		}
	}

	err := s.insert(value, idx)

	if err != nil {
		fmt.Println("Deu erro ai po")
	}

}

func (s *Set) insert(value int, index int) error {

	if s.size >= s.capacity {
		s.Reserve(s.capacity * 2)
	}

	if index < 0 || index >= s.size+1 {
		fmt.Println("index errado po")
	}

	first := s.data[index]
	s.data[index] = value

	for i := index; i < s.size; i++ {

		tmp := s.data[i+1]
		s.data[i+1] = first
		first = tmp
	}

	s.size += 1

	return nil
}

func (s *Set) Erase(value int) {

	idx := s.binarySearch(value)

	if idx == -1 {
		fmt.Println("value not found")
		return
	}

	err := s.erase(idx)

	if err != nil {
		fmt.Println(err)
	}

}

func (s *Set) erase(index int) error {

	if index < 0 || index >= s.size {
		return errors.New("index out of range")
	}

	s.data[index] = 0

	for i := index; i < s.size-1; i++ {
		tmp := s.data[i+1]
		s.data[i+1] = s.data[i]
		s.data[i] = tmp
	}

	s.size -= 1

	return nil
}

func (s Set) binarySearch(value int) int {

	left := 0
	right := s.size - 1

	for left <= right {
		mid := left + (right-left)/2

		if s.data[mid] == value {
			return mid
		}

		if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func (s Set) Contains(value int) bool {

	idx := s.binarySearch(value)

	if idx != -1 {
		return true
	} else {
		return false
	}
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	s := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			s = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				s.Insert(value)
			}
		case "show":
			fmt.Println(s)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			s.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(s.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
