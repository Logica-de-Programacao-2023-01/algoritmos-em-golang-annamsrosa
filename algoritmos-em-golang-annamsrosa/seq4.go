package main

import "fmt"

func main() {
	var (
		num1 float64
		num2 float64
		num3 float64
	)
	fmt.Println("Digite 3 números: ")
	fmt.Scan(&num1, &num2, &num3)
	MP := (2*num1 + 3*num2 + 5*num3) / (2 + 3 + 5)
	fmt.Println("A média ponderada é: ", MP)
}
