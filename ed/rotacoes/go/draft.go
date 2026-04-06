package main

import "fmt"

func main() {

	var qtd_vector, qtd_rotations int

	fmt.Scan(&qtd_vector, &qtd_rotations)

	vector := make([]int, qtd_vector)
	vector_rotated := make([]int, qtd_vector)

	for i := 0; i < qtd_vector; i++ {

		fmt.Scan(&vector[i])

	}

	for i := qtd_vector - 1; i >= 0; i-- {
		value_on_last_idx := vector[i]
		idx_upgraded := (i + qtd_rotations) % qtd_vector

		vector_rotated[idx_upgraded] = value_on_last_idx
	}

	fmt.Print("[ ")
	for _, val := range vector_rotated {
		fmt.Print(val)
		fmt.Print(" ")
	}
	fmt.Println("]")

}
