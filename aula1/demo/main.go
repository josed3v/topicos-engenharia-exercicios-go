package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice() {
	rand.Seed(time.Now().UnixNano())

	min := 1
	max := 8
	var res, php, mhp int
	hp := 50
	hp2 := 35

	fmt.Println("\nRole os Dados apertando Enter até que um dos dois seja derrotado !")
	fmt.Scanln()

	for hp >= 0 && hp2 >= 0 {
		res = rand.Intn(max-min+1) + min
		fmt.Printf("\nA face sorteada foi: %d\n", res)

		php = hp - res
		hp = php
		fmt.Printf("Sua vida desceu para %d !", php)

		mhp = hp2 - res
		hp2 = mhp
		fmt.Printf("\nA vida do adversario desceu para %d! \n", mhp)

		fmt.Scanln()
	}

	if hp >= 0 || hp2 >= 0 {
		fmt.Println("\nAdversario derrotado! ")
	} else {
		fmt.Println("\nGame Over!!! ")
	}
}
func main() {
	var nome string
	var classe int
	fmt.Println("Insira seu nome: ")
	fmt.Scanln(&nome)
	fmt.Println("Bem-vindo à sua nova aventura, Sir", nome)

	fmt.Println("Escolha a sua classe:\n1 - guerreiro\n2 - arqueiro\n3 - feiticeiro")
	fmt.Scanln(&classe)
	switch classe {
	case 1:
		fmt.Println(nome, "O cavaleiro errante está afiando a espada para a missão")

		fmt.Println("Teste sua habilidade contra o javali de treino!")
		rollDice()

		fmt.Println("Agora que voce ja me mostrou seu potencial em batalhas, vamos ao que interessa")
		fmt.Println("Eu preciso saber o seu potencial logico e analitico.")
		fmt.Println("Me responda essa questão:")
		fmt.Println("Qual a criatura que durante a manhã anda sobre 4 patas")
		fmt.Println("á tarde, anda sobre 2 patas")
		fmt.Println("e a noite, sobre 3 patas")
		var res string
		fmt.Scanln(&res)
		if res == "homem" || res == "humano" {
			fmt.Println("A sua resposta esta correta")
			fmt.Println("Voce esta pronto para continuar, agora junte suas coisas e vamos")
			fmt.Println("A sua jornada vai ser ardua mas trata bons frutos...")
			fmt.Println("Ultimamente vem acontecido mais ataques de dragão do que o normal.")
			fmt.Println("Só essa semana fora dizimadas 19 aldeias e a previsão é de mais 2 até o fim do dia.")

		} else {
			fmt.Println("Esqueça tudo que eu te disse até aqui e va para sua casa")
		}
	case 2:
		fmt.Println(nome, "O sagaz arqueiro está preparando suas flechas para a missão")

		fmt.Println("Teste sua habilidade contra o faisão de treino!")
		rollDice()

		fmt.Println("Agora que voce ja me mostrou seu potencial em batalhas, vamos ao que interessa")
		fmt.Println("Eu preciso saber o seu potencial logico e analitico.")
		fmt.Println("Me responda essa questão:")
		fmt.Println("Qual a criatura que durante a manhã anda sobre 4 patas")
		fmt.Println("á tarde, anda sobre 2 patas")
		fmt.Println("e a noite, sobre 3 patas")
		var res string
		fmt.Scanln(&res)
		if res == "homem" || res == "humano" {
			fmt.Println("A sua resposta esta correta")
			fmt.Println("Voce esta pronto para continuar, agora junte suas coisas e vamos")
			fmt.Println("A sua jornada vai ser ardua mas trata bons frutos...")
			fmt.Println("Ultimamente vem acontecido mais ataques de dragão do que o normal.")
			fmt.Println("Só essa semana fora dizimadas 19 aldeias e a previsão é de mais 2 até o fim do dia.")
		} else {
			fmt.Println("Esqueça tudo que eu te disse até aqui e va para sua casa")
		}
	case 3:
		fmt.Println(nome, "O sábio feiticeiro está preparado para a missão")

		fmt.Println("Teste sua habilidade contra o sapo de treino!")
		rollDice()

		fmt.Println("Agora que voce ja me mostrou seu potencial em batalhas, vamos ao que interessa")
		fmt.Println("Eu preciso saber o seu potencial logico e analitico.")
		fmt.Println("Me responda essa questão:")
		fmt.Println("Qual a criatura que durante a manhã anda sobre 4 patas")
		fmt.Println("á tarde, anda sobre 2 patas")
		fmt.Println("e a noite, sobre 3 patas")
		var res string
		fmt.Scanln(&res)
		if res == "homem" || res == "humano" {
			fmt.Println("A sua resposta esta correta")
			fmt.Println("Voce esta pronto para continuar, agora junte suas coisas e vamos")
			fmt.Println("A sua jornada vai ser ardua mas trata bons frutos...")
			fmt.Println("Ultimamente vem acontecido mais ataques de dragão do que o normal.")
			fmt.Println("Só essa semana fora dizimadas 19 aldeias e a previsão é de mais 2 até o fim do dia.")
		} else {
			fmt.Println("Esqueça tudo que eu te disse até aqui e va para sua casa")
		}
	default:
		fmt.Println("Classe inválida!")
	}
}
