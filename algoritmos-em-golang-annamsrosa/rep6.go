package main

import "fmt"

func main() {
	var num float64
	fmt.Println("Digite um número: ")
	fmt.Scan(&num)
	fmt.Println("A tabuada do número digitado é: ")
	for i := 10; i <= 10; i++ {
		tabuada := float64(i) * num
		fmt.Println(tabuada)
	}
}
