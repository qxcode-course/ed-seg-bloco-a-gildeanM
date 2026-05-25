package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func findElementBySet(originIdx int, L int, val int, seq string) bool {

// 	if originIdx >= len(seq) || originIdx < 0 {
// 		return true
// 	}

// 	if int(seq[originIdx]-'0') == val {
// 		return false
// 	}

// 	if L == 0 {
// 		return true
// 	}

// 	return findElementBySet(originIdx-1, L-1, val, seq) && findElementBySet(originIdx+1, L-1, val, seq)

// }

func solve(pos int, seq string, L int) (bool, string) {

	idx := strings.IndexAny(seq[pos:], ".")

	if idx == -1 {
		return true, seq
	}

	idx += pos

	for i := 0; i <= L; i++ {
		if isValid(idx, i, seq, L) {
			seq = seq[:idx] + strconv.Itoa(i) + seq[idx+1:]
			b, res := solve(idx+1, seq, L)
			if b {
				return true, res
			}
			seq = seq[:idx] + "." + seq[idx+1:]
		}
	}

	return false, seq

}

func isValid(idx int, d int, seq string, L int) bool {

	for i := 1; i <= L; i++ {
		if idx-i < 0 {
			break
		}

		if int(seq[idx-i]-'0') == d {
			return false
		}
	}

	for i := 1; i <= L; i++ {
		if idx+i >= len(seq) {
			break
		}

		if int(seq[idx+i]-'0') == d {
			return false
		}
	}

	return true
}

func main() {

	var seq string
	var L int

	fmt.Scan(&seq, &L)
	b, res := solve(0, seq, L)
	if b {
		fmt.Println(res)
	}
}

// idx := strings.IndexAny(seq, ".") // 3
// for i := 0; i <= L; i++ {

// 	if findElementBySet(idx-1, L-1, i, seq) && findElementBySet(idx+1, L-1, i, seq) {
// 		seq = seq[:idx] + strconv.Itoa(i) + seq[idx+1:]
// 		i = 0
// 		idx = strings.IndexAny(seq, ".")
// 	}

// 	if !strings.ContainsAny(seq, ".") {
// 		break
// 	}
// }
