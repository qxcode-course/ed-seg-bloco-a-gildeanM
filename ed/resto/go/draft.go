package main

import "fmt"

func resto(num int) (int, int) {

	if num == 0 {
		return 0, 0
	} else {
		q, r := resto(num / 2)
		if q == 0 && r == 0 {
		} else {
			fmt.Printf("%d %d\n", q, r)
		}
		return num / 2, num % 2
	}

}

func main() {

	var num int

	fmt.Scan(&num)
	fmt.Println(resto(num))

}
