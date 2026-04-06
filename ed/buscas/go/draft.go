package main

import "fmt"

func main() {

	var nquerys, nsearchs int

	fmt.Scan(&nquerys)

	strs := make(map[string]int, nquerys)

	for i := 0; i < nquerys; i++ {
		var query string
		fmt.Scan(&query)
		strs[query]++
	}

	fmt.Scan(&nsearchs)

	// searchs := make([]string, nsearchs)
	searchsnum := make([]int, nsearchs)

	for i := 0; i < nsearchs; i++ {
		var search string
		fmt.Scan(&search)
		searchsnum[i] = strs[search]
	}

	for i, num := range searchsnum {
		fmt.Printf("%d", num)

		if i != nsearchs-1 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}

}
