package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Digite a sua idade: ")
	fmt.Scan(&idade)
	AD := idade * 365
	fmt.Println("Sua idade em dias é: ", AD)
}
