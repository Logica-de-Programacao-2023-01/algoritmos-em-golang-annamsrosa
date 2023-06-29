package main

import "fmt"

func main() {
	var (
		diasTrab    int
		valorDiário float64
	)
	fmt.Println("Digite o número de dias trabalhados: ")
	fmt.Scan(&diasTrab)
	fmt.Println("Digite o valor da diária: ")
	fmt.Scan(&valorDiário)
	S := float64(diasTrab) * valorDiário
	fmt.Println("O seu salário é: ", S)

}
