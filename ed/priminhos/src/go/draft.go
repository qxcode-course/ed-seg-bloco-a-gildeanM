package main

import (
	"fmt"
	"math"
	"strconv"
)

func eh_primo(x int, div int) bool {
	_, _ = x, div

	if x <= 1 {
		return false
	}
	if x <= 3 {
		return true
	}
	if x%2 == 0 || x%3 == 0 {
		return false
	}

	if div > int(math.Sqrt(float64(x)))+1 {
		return true
	} else {
		if x%div == 0 || x%(div+2) == 0 {
			return false
		}
		return eh_primo(x, div+6)
	}

}

func getPrimes(p int, n int, primes []int) []int {

	if eh_primo(p, 5) {
		// fmt.Println(strconv.Itoa(p) + "eh primo")
		primes = append(primes, p)
	}

	if len(primes) < n {
		return getPrimes(p+1, n, primes)
	}

	return primes
}

func main() {

	var n int

	fmt.Scan(&n)

	primes := make([]int, 0)

	primes = getPrimes(2, n, primes)

	str := "["
	for i := 0; i < n; i++ {
		str += strconv.Itoa(primes[i])
		if i != n-1 {
			str += ", "
		}
	}
	str += "]"

	fmt.Println(str)

}
