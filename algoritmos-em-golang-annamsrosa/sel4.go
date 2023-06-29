package main

import "fmt"

func main() {
	var (
		altura   float64
		sexo     string
		situação string
	)
	fmt.Println("Digite a altura: ")
	fmt.Scan(&altura)
	fmt.Println("Digite o sexo (Feminino ou Masculino")
	fmt.Scan(&sexo)

	pesoIdeal := 0.0
	if sexo == "F" {
		pesoIdeal = 62.1*altura - 44.7
	} else if sexo == "M" {
		pesoIdeal = 72.7*altura - 58
	}
	if pesoIdeal > 0 {
		situação = "acima do peso ideal"
	} else if pesoIdeal > 0 {
		situação = "abaixo do peso ideal"
	} else {
		situação = "no peso ideal"
	}
	fmt.Println("A pessoa se encontra ", situação)
}
