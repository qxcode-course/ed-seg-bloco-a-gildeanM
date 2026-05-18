package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func (s Set) binarySearch(value int) int {

// 	left := 0
// 	right := s.size - 1

// 	for left <= right {
// 		mid := left + (right-left)/2

// 		if s.data[mid] == value {
// 			return mid
// 		}

// 		if s.data[mid] < value {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	return -1
// }

func BetterSearch(slice []int, value int) (bool, int) {
	_, _ = slice, value

	left := 0
	right := len(slice) - 1

	for left <= right {
		mid := left + (right-left)/2

		if slice[mid] == value {
			return true, mid
		}

		if slice[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	probable_position := 0
	for i := 0; i < len(slice)-1; i++ {

		if value > slice[i] && value < slice[i+1] {
			probable_position = i + 1
			break
		}

		if value > slice[len(slice)-1] {
			probable_position = len(slice)
			break
		}
	}
	return false, probable_position

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
