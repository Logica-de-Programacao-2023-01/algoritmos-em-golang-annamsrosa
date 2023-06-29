package main

import "fmt"

func main() {
	var num int
	fmt.Println("Digite um número: ")
	fmt.Scan(&num)
	S := num + 1
	A := num - 1
	fmt.Println("O sucessor e o antecessor desse númuro são: ", S, A)
}
