package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {

	slices.SortFunc(vet, func(a, b int) int {
		return cmp.Compare(math.Abs(float64(a)), math.Abs(float64(b)))
	})

	max := int(math.Abs(float64(vet[len(vet)-1]))) + 1

	occ := make([]int, max)

	for _, val := range vet {
		idx := int(math.Abs(float64(val)))
		occ[idx]++
	}

	pairs := make([]Pair, 0, max)

	for idx := range occ {
		if occ[idx] > 0 {
			pairs = append(pairs, Pair{One: idx, Two: occ[idx]})
		}
	}

	return pairs
}

func teams(vet []int) []Pair {

	pairs := make([]Pair, 0, len(vet))

	for _, val := range vet {
		if len(pairs) > 0 && pairs[len(pairs)-1].One == val {
			pairs[len(pairs)-1].Two++
		} else {
			pairs = append(pairs, Pair{One: val, Two: 1})
		}
	}

	return pairs
}

func mnext(vet []int) []int {
	mapaprox := make([]int, len(vet))
	n := len(vet)

	for i := range vet {

		if i == 0 {
			if n == 1 && vet[i] > 0 {
				continue
			} else if n > 1 {
				if vet[i] > 0 && vet[i+1] > 0 {
					continue
				}
			}
		} else if i == (n - 1) {
			if vet[i] < 0 || vet[i] > 0 && vet[i-1] > 0 {
				continue
			}
		} else {
			if vet[i] > 0 && vet[i+1] > 0 && vet[i-1] > 0 {
				continue
			}
		}

		if vet[i] < 0 {
			continue
		}

		mapaprox[i]++
	}

	return mapaprox
}

func alone(vet []int) []int {
	mapaprox := make([]int, len(vet))
	n := len(vet)

	for i := range vet {

		if i == 0 {
			if n == 1 && vet[i] < 0 {
				continue
			} else if n > 1 {
				if vet[i] > 0 && vet[i+1] < 0 {
					continue
				}
			}
		} else if i == (n - 1) {
			if vet[i] < 0 || vet[i] > 0 && vet[i-1] < 0 {
				continue
			}
		} else {
			if vet[i] > 0 && (vet[i+1] < 0 || vet[i-1] < 0) {
				continue
			}
		}

		if vet[i] < 0 {
			continue
		}

		mapaprox[i]++
	}

	return mapaprox
}

func couple(vet []int) int {

	qtdPersons := make(map[int]int, len(vet))

	for _, val := range vet {
		qtdPersons[val] += 1
	}

	visited := make(map[int]bool, len(vet))
	couples := 0

	for _, val := range vet {

		if !visited[val] {

			man := qtdPersons[val]
			woman := qtdPersons[-val]

			if man == woman {
				couples += man
			} else if man > woman {
				couples += woman
			} else if woman > man {
				couples += man
			}

			visited[val] = true
			visited[-val] = true
		}
	}

	return couples
}

func hasSubseq(vet []int, seq []int, pos int) bool {

	if (len(vet) - pos) < len(seq) {
		return false
	}

	for i := 0; i < len(seq); i++ {
		if vet[pos] != seq[i] {
			return false
		}
		pos++
	}

	return true
}

func subseq(vet []int, seq []int) int {
	_ = vet
	_ = seq

	for i := 0; i < len(vet); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}

	return -1
}

func erase(vet []int, posList []int) []int {

	erasedvet := make([]int, 0, len(vet))
	mapspos := make(map[int]bool, len(posList))

	for _, val := range posList { //n * 1
		mapspos[val] = true
	}

	for idx, val := range vet { //m * 1
		if !mapspos[idx] {
			erasedvet = append(erasedvet, val)
		}
	}

	return erasedvet
}

func clear(vet []int, value int) []int {

	clearedByValueVet := make([]int, 0, len(vet))

	for _, val := range vet {
		if val != value {
			clearedByValueVet = append(clearedByValueVet, val)
		}
	}

	return clearedByValueVet
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
