package main

import "fmt"

func main() {
	var salário float64
	fmt.Println("Digite o seu salãrio: ")
	fmt.Scan(&salário)
	if salário < 1000 {
		NS := salário * 1.1
		fmt.Println("O seu salário com aumento é: ", NS)
	} else {
		NS := salário * 1.05
		fmt.Println("O seu salário com aumento é: ", NS)
	}
}
