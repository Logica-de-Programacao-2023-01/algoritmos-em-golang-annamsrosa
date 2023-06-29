package main

import "fmt"

func main() {
	var salário float64
	fmt.Println("Digite o seu salário: ")
	fmt.Scan(&salário)
	NS := salário * 1.15
	fmt.Println("O seu novo salário com um aumento de 15% é ", NS)
}
