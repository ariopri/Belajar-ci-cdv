package main

import "fmt"

/*
	Recursive function -> function yang memanggil dirinya sendiri
	contoh :
	Kasusnya itu faktorial
	Di kasus kali ini kita akan buat menjadi  bagian
	1. dengan forloop
	2. dengan recursive function
*/

func ForLooping(parameter int) int {
	result := 1
	for i := parameter; i > 0; i-- {
		result *= i
	}
	return result
}

func ForLoop(n int) int {
	var result int = 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}

func main() {
	fmt.Println(ForLoop(5))
	fmt.Println(ForLooping(6))
	fmt.Println(RecursiveFunction(4))
}

// ini dengan Recursive Function
func RecursiveFunction(nilai int) int {
	if nilai == 1 {
		return 1
	} else {
		return nilai * RecursiveFunction(nilai-1)
	}
	/*
		penjelasan:
		dia tergantung (nilai int)
		jika gw masukin nilainya 4, maka nilai pada
		return nilai * RecursiveFunction(nilai-1) akan menjadi 4
		4 * RecursiveFunction(3)
		nah dia balik nih RecursiveFunction(3) ke RecursiveFunction(nilai int)
		habis itu dia jadi 3 * RecursiveFunction(2)
		nah begtu selanjunya
		jika dia 1 maka dia akan mereturn 1 dan keluar dari looping
	*/
}
