package main

import "fmt"

func newInt() *int {
	return new(int)
}

func main() {

	p := newInt()
	q := newInt()

	fmt.Println(p == q) // Retornará falso pq não apontam para lugares diferentes na memória

	fmt.Println(*p == *q) // Retorna true pq acessamos o valores dos ponteiros e comparamos, como o valor default é 0
}
