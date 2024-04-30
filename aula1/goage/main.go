package main

import (
	"fmt"
)

func isLegalAge(age int) bool {
	return age >= 18
}

func main() {
	var age int
	fmt.Print("Por favor, insira sua idade: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println("Erro ao ler a idade:", err)
		return
	}

	if isLegalAge(age) {
		fmt.Println("É maior de idade.")
	} else {
		fmt.Println("É menor de idade.")
	}
}
