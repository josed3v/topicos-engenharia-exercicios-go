package main

import "fmt"

func main() {
	lista := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	m := make(map[int]int)

	for i, v := range lista {
		m[i] = v
	}

	fmt.Println("Mapa:", m)
}
