package main

import "fmt"

type Person struct {
	Name, School string
	Kelas        int
}

func main() {
	//struct sebuah template data yang digunakan untuk mengabungkan tipe data
	// untuk mengabungkan beberapa tipe data makanya coocknya struct
	var jujutsu Person
	jujutsu.Name = "Itadori"
	jujutsu.School = "SMK Jujutsu"
	jujutsu.Kelas = 1

	fmt.Println(jujutsu)
	fmt.Println(jujutsu.Name)
	fmt.Println(jujutsu.School)
	fmt.Println(jujutsu.Kelas)

	kaisen := Person{
		Name:   "Gojo Satoru",
		School: "SMK Jujutsu",
		Kelas:  100,
	}
	fmt.Println(kaisen)
	fmt.Println(kaisen.Name)
	fmt.Println(kaisen.School)
}
