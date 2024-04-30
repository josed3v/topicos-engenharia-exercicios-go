package main

import (
	"errors"
	"fmt"
)

func countString(s string) (int, error) {
	if s == "" {
		return 0, errors.New("a string está vazia")
	}
	return len(s), nil
}

func main() {
	str := "Hello world!"
	fmt.Println(str)
	tamanho, err := countString(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("O tamanho da string é %d\n", tamanho)
	}
}
