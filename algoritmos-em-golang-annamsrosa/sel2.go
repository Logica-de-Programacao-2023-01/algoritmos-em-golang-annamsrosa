package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
		num3 int
	)
	fmt.Println("Digite 3 números inteiros: ")
	fmt.Scan(&num1, &num2, &num3)
	menor := num1
	if num2 < menor {
		menor = num2
	}
	if num3 < menor {
		menor = num3
	}
	fmt.Println("O menor númeroe: ", menor)
}
