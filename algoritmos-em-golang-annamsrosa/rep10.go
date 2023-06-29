package main

import (
	"fmt"
)

func main() {
	var num int
	maior := 0
	menor := 0
	for {
		fmt.Println("Digite um número: ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num > maior {
			maior = num
		}
		if num < maior {
			menor = num
		}
	}
	fmt.Println("O maior número é: ", maior, ", e o menor é: ", menor)
}
