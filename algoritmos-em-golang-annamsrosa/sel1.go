package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
	)
	fmt.Println("Digite 2 números interios: ")
	fmt.Scan(&num1, &num2)
	maior := num1
	if num2 > maior {
		maior = num2
	}
	fmt.Println("O maior número é: ", maior)
}
