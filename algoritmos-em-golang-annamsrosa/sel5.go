package main

import "fmt"

func main() {
	var n int
	fmt.Println("Diigite um número inteiro: ")
	fmt.Scan(&n)
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("O número é múltiplo de 3 e 5 ao mesmo tempo ")
	} else {
		fmt.Println("O número não é múltiplo de 3 e 5 ao mesmo tempo ")
	}
}
