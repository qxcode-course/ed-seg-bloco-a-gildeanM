package main

import "fmt"

const HORARIO = -1
const ANTI_HORARIO = 1

func calc_distance(from int, to int, direction int) int {
	isHorario := 0
	if direction == HORARIO {
		isHorario = 1
	}
	distance := 0
	if to > from {
		if isHorario == 1 {
			distance = 16 - (to - from)
		} else {
			distance = to - from
		}
	} else {
		if isHorario == 1 {
			distance = from - to
		} else {
			distance = 16 - (from - to)
		}
	}
	return distance
}

func main() {
	var h, p, f, d int
	var res byte = 'S'

	fmt.Scanf("%d %d %d %d", &h, &p, &f, &d)

	distanceFromFToH := calc_distance(f, h, d)
	distanceFromPToH := calc_distance(p, h, d)

	if distanceFromFToH > distanceFromPToH {
		res = 'N'
	}

	fmt.Printf("%c\n", res)
}
