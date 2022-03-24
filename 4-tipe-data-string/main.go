package main

import "fmt"

func newFunction1(nama, sekolah, kota string) {
	fmt.Println("Biodata")
	fmt.Println("1. Nama: ", nama)
	fmt.Println("2. Sekolah: ", sekolah)
	fmt.Println("3. Kota Asal:", kota)
}

func newFunction() {
	var (
		nama    string
		sekolah string
		kota    string
	)

	fmt.Printf("Masukan namanya disini: %s", nama)
	fmt.Scanf("%s", &nama)
	fmt.Printf("Masukan nama sekolahnya  disini: %s", sekolah)
	fmt.Scanf("%s %s", &sekolah)
	fmt.Printf("masukan nama asal kota disini: %s", kota)
	fmt.Scanf("%s", &kota)

	newFunction1(nama, sekolah, kota)
}

func main() {
	// atau yg cepat nama, sekolah, kota string
	newFunction()
}
