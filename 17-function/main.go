package main

import "fmt"

func goodmorning() {
	fmt.Println("Selamat pagi")
}

func main() {
	goodmorning()
	sayHello(" Itadori Yuji", "Gojo Satoru")
	// result := getHello("Itadori Yuji") // string ini panjanganya 12
	// fmt.Println(result)
	// atau
	fmt.Println(getHello(10))
	//ReturningMultipleValues
	no, angka, nama := ReturningMultipleValues()
	fmt.Println(no, angka, nama)
	//NamedReturnValues
	_, nama1, nama2 := NamedReturnValues()
	fmt.Println(nama1, nama2)

}

/* function parameter
hal ini tidak wajib
*/
func sayHello(nama1, nama2 string) {
	fmt.Println("Hello", nama1, nama2)
}

/*
	 function return value
	perhatikan baik2 coy janga sampe lupa dan gak bisa
	func getHello(name string) string
	name string -> parameter
	string -> nah yang ini nih yang akan kita return
	atau keluarkan hasilnya
*/

func getHello(name int) bool {
	// len(name) berarti dia akan mereturn hasilnya sebagai int
	// harus mengunakan fungsi len()
	if name == 1 {
		return true
	} else {
		return false
	}
}

func ReturningMultipleValues() (string, int, string) {
	/*
		ReturningMultipleValues
		-> mengembalikan data lebih dari satu
		cth getfullname() (string, string)

	*/
	return "no.", 1, "Itadori"

}

func NamedReturnValues() (no int, nama1 string, nama2 string) {
	/*
		func NamedReturnValues(nama1, nama2, nama3 string)
		bisa langsung bikin varible di func secara lansugn

	*/
	no = 1
	nama1 = "Itadori"
	nama2 = "sukuna"
	return
}
