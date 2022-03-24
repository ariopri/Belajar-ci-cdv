package main

import "fmt"

/*
	closure adalah kemampuan sebuah function berinteraksi dengan
	data2 di sekitarnya dengan scope yang sama

*/

func main() {
	// contoh lagi nih
	name := "Itadori Yuji"
	counter := 0
	increment := func() int {
		name = "Sukuna"
		// dia bakanlan ngerubah nih namenya jadi sukuna
		counter++
		return counter
	}
	increment()
	increment()
	increment()
	// hasilnya 3 karena gw painggil increment 3 kali
	// dia ngerubah nih counter di luarnya, nah makanya hati2 menggunakan closure
	fmt.Println("Counter : ", counter)
	fmt.Println("Name : ", name)
}
