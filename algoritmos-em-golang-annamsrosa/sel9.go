package main

import "fmt"

func main() {
	var (
		num1 float64
		num2 float64
		num3 float64
	)
	fmt.Println("Digite 3 números reais: ")
	fmt.Scan(&num1, &num2, &num3)

	if num1 <= num2 && num1 <= num3 {
		fmt.Println(num1, " ")
		if num2 <= num3 {
			fmt.Println(num2, num3)
		} else {
			fmt.Println(num3, num2)
		}
	} else if num2 <= num1 && num2 <= num3 {
		fmt.Println(num2, " ")
		if num1 <= num3 {
			fmt.Println(num1, num3)
		} else {
			fmt.Println(num3, num1)
		}
	} else {
		fmt.Print(num3, " ")
		if num1 <= num2 {
			fmt.Println(num1, num2)
		} else {
			fmt.Println(num2, num1)
		}
	}
}
