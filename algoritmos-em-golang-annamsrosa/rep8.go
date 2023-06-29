package main

import "fmt"

func main() {
	var num int
	fmt.Println("Digite um número inteiro e positivo: ")
	fmt.Scan(&num)
	for i := 1; i <= 100000; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
}
