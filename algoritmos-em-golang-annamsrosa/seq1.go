package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
		num3 int
	)
	fmt.Println("Digite 3 números inteiros")
	fmt.Scan(&num1, &num2, &num3)
	soma := num1 + num2 + num3
	fmt.Println("A soma deles é: ", soma)
}
