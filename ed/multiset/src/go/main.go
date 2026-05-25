package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *MultiSet) Reserve(newCapacity int) {

	if newCapacity <= s.capacity {
		return
	}

	newSet := NewMultiSet(newCapacity)

	newSet.size = s.size

	for i := 0; i < s.size; i++ {
		newSet.data[i] = s.data[i]
	}

	*s = *newSet

}

func (s *MultiSet) Insert(value int) {

	idx := s.size

	if s.Contains(value) {
		idx = s.binarySearch(value)
	}

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

func (s *MultiSet) insert(value int, index int) error {

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

func (s *MultiSet) Clear() {
	for i := s.size - 1; i >= 0; i-- {
		s.erase(i)
	}
}

func (s *MultiSet) Erase(value int) {

	if !s.Contains(value) {
		fmt.Println("value not found")
		return
	}

	idx := s.binarySearch(value)

	err := s.erase(idx)

	if err != nil {
		fmt.Println(err)
	}

}

func (s *MultiSet) erase(index int) error {

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

func (s MultiSet) binarySearch(value int) int {

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

func (s MultiSet) Contains(value int) bool {

	idx := s.binarySearch(value)

	if idx != -1 {
		return true
	} else {
		return false
	}
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (s MultiSet) String() string {

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

func (s MultiSet) Unique() int {

	uniques := NewMultiSet(s.size)

	for i := 0; i < s.size; i++ {
		value := s.data[i]
		if !uniques.Contains(value) {
			uniques.Insert(value)
		}
	}

	return uniques.size
}

func (s MultiSet) Count(value int) int {

	count := 0

	for _, val := range s.data {
		if val == value {
			count++
		}
	}

	return count

}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			value, _ := strconv.Atoi(args[1])
			ms.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
