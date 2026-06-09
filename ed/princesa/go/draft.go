package main

import (
	"fmt"
	"strconv"
	"strings"
)

func procurar_vivo(deads map[int]bool, pos int, size int) int {

	for {
		if deads[pos] {
			pos = (pos + 1) % size
		} else {
			return pos
		}
	}
}

func princesa(elementos []int, has_sword int) {

	len := len(elementos)
	n := len
	deads := make(map[int]bool, len)

	for n > 1 {
		fmt.Println(String(elementos, has_sword, deads))
		pos := procurar_vivo(deads, has_sword, len)
		deads[pos] = true
		has_sword = procurar_vivo(deads, pos, len)
		n--
	}

}

func String(elementos []int, posHasSword int, deads map[int]bool) string {
	var b strings.Builder

	b.WriteString("[ ")
	for idx, val := range elementos {
		if deads[idx] {
			continue
		}
		b.WriteString(strconv.Itoa(val))
		if idx == posHasSword {
			b.WriteByte('>')
		}
		b.WriteByte(' ')
	}
	b.WriteString("]")

	return b.String()
}

func main() {
	var qtd_elementos, has_sword int

	fmt.Scan(&qtd_elementos, &has_sword)

	elementos := make([]int, qtd_elementos)

	for i := 0; i < qtd_elementos; i++ {
		elementos[i] = i + 1
	}

	princesa(elementos, has_sword-1)

}
