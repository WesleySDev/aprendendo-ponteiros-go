package main

import "fmt"

func dobraValor(numero *int) {
	*numero = *numero * 2
}

func trocaValor(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	numero := 2

	dobraValor(&numero)

	fmt.Println(numero)

	a := 10
	b := 20

	trocaValor(&a, &b)

	fmt.Println("a:", a)
	fmt.Println("b:", b)

}
