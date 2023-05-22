package main

import (
	"fmt"
	"sort"
)

func main() {
	Jogo()
}

func Jogo() {

	numero := false

	random := 2 //rand.Intn(100) + 1

	n_Tries := make(map[int]int)
	tries := 0

	var choose int

	fmt.Println("Bem vindo ao jogo de adivinhação!")
	for !numero {
		fmt.Print("Insira um número entre 1 e 100: ")
		fmt.Scan(&choose)
		tries++
		if choose > 100 || choose <= 0 {
			fmt.Println("Número inválido.")
		} else if choose < random {
			fmt.Println("O número é maior que", choose, ":")
		} else if choose > random {
			fmt.Println("O número é menor que", choose, ":")
		} else {
			fmt.Println("Parabéns, Você acertou!")
			fmt.Printf("Você utilizou %d tentativa(s).\n", tries)
			numero = true
		}

	}
	fmt.Println("Bora jogar de novo? (s/n)")

	var jogarnovamente string
	fmt.Scan(&jogarnovamente)

	if jogarnovamente != "s" && jogarnovamente != "S" {
		fmt.Println("Obrigado por jogar.")
		n_Tries_ord := make([]int, 0, len(n_Tries))

		for n_index := range n_Tries {
			n_Tries_ord = append(n_Tries_ord, n_index)

		}

		sort.Ints(n_Tries_ord)

		for _, tries := range n_Tries_ord {
			fmt.Printf("Você utilizou %d tentativas na %dº tentativa \n", n_Tries[tries], n_Tries[tries])
		}
		return
	}
	Jogo()
}
