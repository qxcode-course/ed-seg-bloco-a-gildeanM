package main

import (
	"fmt"
)

func checkSum(set []int, totalSum int) bool {

	if totalSum < 0 {
		return false
	}

	if totalSum == 0 {
		return true
	}

	if len(set) == 0 {
		return false
	}

	// fmt.Println(set, totalSum)
	return checkSum(set[1:], totalSum-set[0]) || checkSum(set[1:], totalSum)
}

func main() {

	var qtd_subset, total_sum int

	fmt.Scan(&qtd_subset, &total_sum)

	set := make([]int, 0, qtd_subset)

	for i := 0; i < qtd_subset; i++ {
		var value int
		fmt.Scan(&value)
		set = append(set, value)
	}

	fmt.Println(checkSum(set, total_sum))

}
