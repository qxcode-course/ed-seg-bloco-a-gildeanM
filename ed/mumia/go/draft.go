package main

import "fmt"

const NAME_LEN = 25
const CLASSIFICATION_LEN = 8

func main() {
	var name string
	var age int
	fmt.Scan(&name)
	fmt.Scan(&age)

	var classification string

	if age < 12 {
		classification = "crianca"
	} else if age < 18 {
		classification = "jovem"
	} else if age < 65 {
		classification = "adulto"
	} else if age < 1000 {
		classification = "idoso"
	} else {
		classification = "mumia"
	}

	fmt.Printf("%s eh %s\n", name, classification)
}
