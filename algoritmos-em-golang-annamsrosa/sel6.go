package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
	)
	fmt.Println("Digite 2 números inteiros: ")
	fmt.Scan(&num1, &num2)
	if num1 >= 0 && num2 >= 0 {
		M := num1 * num2
		fmt.Println("A multiplicação é: ", M)
	} else {
		S := num1 + num2
		fmt.Println("A soma é: ", S)
	}
}
