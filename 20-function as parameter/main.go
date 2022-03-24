package main

import "fmt"

func Running() {
	SayHellowithFilter("Yuji", SpamFilter)
	SayHellowithFilter("Anjing", SpamFilter)
}
func main() {
	// function as parameter
	/*
		function bisa digunakan sebagai parameter untuk function lain
		PERCAYALAH INI SANGAT SUSAH DAN MENYEBALKAN
	*/
	// kita akan mengfilter kata2 kasar
	//cara 1
	SayHellowithFilter("Yuji", SpamFilter) //-> SpamFilter merupakan nama functionya
	// cara 2
	filter := SpamFilter
	SayHellowithFilter("Anjing", filter)
	//cara 3
	Running()

}

//TYPE Declaration
/*
	solusi jila parameter functionnya terlalu panjanganya
	contoh :
	func SayHellowithFilter(name string, filter func(string, string, string, int, int) string) {
	oleh karena itu dibutuhkan type declaration
*/
//contoh type Declaration
type Filter func(string) string

func SayHellowithFilter(name string, filter Filter /*func(string) string -> Disigkat jadi Filter*/) {
	/*
		Penjelasan:
		SayHellowithFilter adalah function yang memiliki 2 parameter
		yaitu (name string, filter func(string) string) -> adalah sebuah
		function yang memiliki 2 parameter yaitu sebagai string dan satu lagi
		parameter yang dideklarasikan sebagai function
	*/
	NamedFiltered := filter(name) // kita akan meng-filter kata2 kasar
	// artinya sebelum di print di bawah makan kita filter terlebih dahulu
	fmt.Println("Hello ", NamedFiltered)
	/*
		filter(name) -> karena filter func(string) parameterinya sebagai string
		contohnya memasukan name
	*/
}
func SpamFilter(name string) string {
	if name == "Anjing" {
		return "*****"
	} else {
		return name
	}
}
