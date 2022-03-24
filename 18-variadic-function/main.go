package main

import "fmt"

func main() {
	//biasa
	total := sumAll(1, 2, 3, 4)
	fmt.Println(total)
	//lebih singkat
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// slice parameter
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sumAll(slice...))

}

func sumAll(numbers ...int) int {
	/*
		func sumAll(number ...int)int
		number ...int -> parameter yang bersifat variable argumen
		-> artinya bisa memanggi sumAll bisa memasukin int lebih dari 1
		maka dia akan berubah menjadi array
		CUMA BISA SATU PARAMETER ITUPUN DI BAGIAN BELAKANG
		atau bisa juga dibikin seperti
		-> func sumAll (number []int) int -> masih abu2ju
	*/
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total

}
