package main

import "fmt"

func main() {
	var (
		nama string
		asal string
	)
	fmt.Printf("masukan nama: %s", nama)
	fmt.Scanf("%s", &nama)
	fmt.Printf("masukan kota asal: %s", asal)
	fmt.Scanf("%s", &asal)
	if nama == "yuji" {
		newFunction()
	} else {

		fmt.Println("Salah masukan input")

	}
}

func newFunction() {
	var (
		size, i, j, k, l int
	)
	fmt.Printf("masukan angka :%.0d", size)
	fmt.Scanf("%d", &size)
	for i = 0; i <= size; i++ {
		for j = size; j >= i; j-- {
			fmt.Printf(" ")
		}
		for k = 0; k <= i; k++ {
			fmt.Printf("*")
		}
		for l = 0; l < i; l++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}
