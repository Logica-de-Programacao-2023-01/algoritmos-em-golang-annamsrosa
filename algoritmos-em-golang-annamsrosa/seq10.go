package main

import "fmt"

func main() {
	var peso float64
	fmt.Println("Digite o seu peso em kg: ")
	fmt.Scan(&peso)
	PL := peso * 2.20462
	fmt.Println("O seu peso em libras é: ", PL)
}
