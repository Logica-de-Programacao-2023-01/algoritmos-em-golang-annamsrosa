package main

import "fmt"

func main() {
	var (
		peso   float64
		altura float64
	)
	fmt.Println("Digite o seu peso ")
	fmt.Scan(&peso)
	fmt.Println("Digite a sua altura ")
	fmt.Scan(&altura)
	IMC := peso / (altura * altura)
	fmt.Println("O seu IMC é ", IMC)
}
