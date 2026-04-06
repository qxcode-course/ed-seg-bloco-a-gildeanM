package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	men := make([]int, 0)

	for _, man := range vet {
		if man > 0 {
			men = append(men, man)
		}
	}
	return men
}

func getCalmWomen(vet []int) []int {
	women := make([]int, 0)

	for _, woman := range vet {
		if woman < 0 && -(woman) < 10 {
			women = append(women, woman)
		}
	}
	return women
}

func sortVet(vet []int) []int {
	slices.Sort(vet)
	return vet
}

func sortStress(vet []int) []int {

	slices.SortFunc(vet, func(i, j int) int {
		return int(math.Abs(float64(i)) - math.Abs(float64(j)))
	})

	return vet
}

func reverse(vet []int) []int {

	reversed := slices.Clone(vet)
	slices.Reverse(reversed)

	return reversed
}

func unique(vet []int) []int {

	already_read := make(map[int]bool)
	uniques := make([]int, 0)

	for _, value := range vet {
		if !already_read[value] {
			uniques = append(uniques, value)
			already_read[value] = true
		}
	}

	return uniques
}

func repeated(vet []int) []int {

	non_repeated := make(map[int]bool)
	repeated := make([]int, 0)

	for _, value := range vet {

		if non_repeated[value] {
			repeated = append(repeated, value)
		} else {
			non_repeated[value] = true
		}
	}

	return repeated
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
