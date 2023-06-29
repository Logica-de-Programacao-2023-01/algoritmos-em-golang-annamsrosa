package main

import "fmt"

func main() {
	var preço float64
	fmt.Println("Digite o preço do produto: ")
	fmt.Scan(&preço)
	PD := preço * 0.9
	fmt.Println("O preço do produto com um desconto de 10% é: ", PD)

}
