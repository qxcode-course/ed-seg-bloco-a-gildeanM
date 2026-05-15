package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v Vector) Status() string {
	status := fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)

	return status
}

func (v Vector) Capacity() int {
	return v.capacity
}

func (v Vector) Size() int {
	return v.size
}

func (v Vector) String() string {

	var vectorString strings.Builder
	vectorString.WriteByte('[')
	for i := 0; i < v.size; i++ {
		vectorString.WriteString(strconv.Itoa(v.data[i]))
		if i != v.size-1 {
			vectorString.WriteString(", ")
		}
	}
	vectorString.WriteByte(']')

	return vectorString.String()
}

func (v Vector) At(index int) (int, error) {

	if index < 0 || index >= v.Size() || index >= v.Capacity() {
		return 0, errors.New("index out of range")
	}

	return v.data[index], nil
}

func (v Vector) IndexOf(value int) int {

	for idx, valAtVector := range v.data {
		if valAtVector == value {
			return idx
		}
	}

	return -1
}

func (v Vector) Contains(value int) bool {

	for _, valAtVector := range v.data {
		if valAtVector == value {
			return true
		}
	}

	return false
}

func (v *Vector) Set(index int, value int) error {

	if index < 0 || index >= v.Size() || index >= v.Capacity() {
		return errors.New("index out of range")
	}

	v.data[index] = value
	return nil
}

func (v *Vector) Reserve(newCapacity int) {

	if newCapacity > v.capacity {
		vNewCapacity := NewVector(newCapacity)
		vNewCapacity.size = v.size

		for i := 0; i < v.size; i++ {
			vNewCapacity.data[i] = v.data[i]
		}

		*v = *vNewCapacity
	}

}

func (v *Vector) PushBack(value int) {
	if v.Size() >= v.Capacity() {
		v.Reserve(v.Capacity() * 2)
	}

	v.data[v.Size()] = value

	v.size += 1

}

func (v *Vector) PopBack() (int, error) {
	lastIdx := v.Size() - 1

	if lastIdx < 0 {
		return 0, errors.New("vector is empty")
	}

	lastValue := v.data[lastIdx]

	v.data[lastIdx] = 0
	v.size -= 1

	return lastValue, nil
}

func (v *Vector) Insert(index int, value int) error {

	v.PushBack(value)

	size := v.size - 1

	for i := 0; i < size; i++ {

		if size-i == index {
			break
		}

		tmp := v.data[size-(i+1)]
		v.data[size-(i+1)] = v.data[size-i]
		v.data[size-i] = tmp

	}

	return nil
}

func (v *Vector) Erase(index int) error {

	if index < 0 || index >= v.Size() || index >= v.Capacity() {
		return errors.New("index out of range")
	}

	v.data[index] = 0

	for ; index < v.Size()-1; index++ {
		tmp := v.data[index+1]
		v.data[index+1] = v.data[index]
		v.data[index] = tmp
	}

	v.size -= 1

	return nil
}

func (v Vector) Slice(start int, end int) *Vector {

	if end < 0 {
		end = v.size + end
	}

	if start < 0 || start >= v.size || end < 0 || end >= v.size {
		return NewVector(0)
	}

	if end < start {
		return NewVector(0)
	}

	capacity := end - start

	slice := NewVector(capacity)

	for i := start; i < end; i++ {

		slice.PushBack(v.data[i])
	}

	return slice
}

func (v *Vector) Clear() {

	for i := v.size; i >= 0; i-- {
		v.PopBack()
	}
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
