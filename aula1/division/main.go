package main

import (
	"fmt"
)

func divideNumbers(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("erro: divisão por zero")
	}
	return float64(a) / float64(b), nil
}

func main() {
	lista1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	lista2 := []int{2, 0, 5, 3, 0, 10, 4, 8, 2, 1}

	for i := 0; i < len(lista1); i++ {
		result, err := divideNumbers(lista1[i], lista2[i])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d / %d = %.2f\n", lista1[i], lista2[i], result)
		}
	}
}
