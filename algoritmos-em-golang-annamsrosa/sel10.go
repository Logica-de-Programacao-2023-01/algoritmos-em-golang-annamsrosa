package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Digite a idade: ")
	fmt.Scan(&idade)

	if idade <= 9 {
		fmt.Println("Classificação: mirim ")
	} else if idade >= 10 && idade <= 13 {
		fmt.Println("Classificação: infantil ")
	} else if idade >= 14 && idade <= 17 {
		fmt.Println("Classificação: juvenil ")
	} else {
		fmt.Println("Classificação: adulto")
	}
}
