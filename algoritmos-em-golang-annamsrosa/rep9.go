package main

import (
	"fmt"
)

func main() {
	var (
		num   int
		quant int
		soma  int
	)
	for {
		fmt.Println("Digite um númeor: ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		quant += 1
		soma += num
	}
	if quant > 0 {
		M := soma / quant
		fmt.Println("A média dos números é: ", float64(M))
	} else {
		fmt.Println("Nenhum número informado ")
	}
}
