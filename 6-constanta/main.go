package main

import "fmt"

func main() {
	//constanta boyy
	const nama = "yuji" //tipe data gk wajib sih kalau ada valuenya
	// kalau yg ini kan kgk
	fmt.Println("Masuka codename :", nama)
	fmt.Println("maka namanya adalah :", nama)
	/*
		nama = "madara";
		karena constanta tidak boleh diuah2 atau diganti
	*/
	cthConstanta()

}
func cthConstanta() {
	school, kelas := newFunction1()
	newFunction(school, kelas)
}

func newFunction1() (string, int) {
	const school = "Smk Jujutsu Tokyo"
	const kelas int = 1
	return school, kelas
}

func newFunction(school string, kelas int) {
	fmt.Println("sekolah :", school)
	fmt.Println("kelas", kelas)
}
