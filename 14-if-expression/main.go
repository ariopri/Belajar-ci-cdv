package main

import "fmt"

func main() {
	ifExpression()
	ifExpression2()
	ifExpression3()
	ifExpression4()
}

func ifExpression() {
	if a := 1; a == 1 {
		println("a is 1")
	} else {
		println("a is not 1")
	}
}
func ifExpression2() {
	name := "madara"

	if name == "madara" {
		println("name is madara")
	} else {
		println("name is not madara")
	}
}

func ifExpression3() {
	var nama string
	fmt.Printf("Masukan namanya disini: %s", nama)
	fmt.Scanf("%s", &nama)
	fmt.Println(nama)

	if nama == "madara" {
		println(true)
	} else if nama == "gojo" {
		println(true)
	} else if nama == "yuji" {
		println(true)
	} else {
		println(false)
	}
}
func ifExpression4() {
	var (
		angka1 int
		angka2 int
	)
	fmt.Printf("Masukan angka 1: %.0d", angka1)
	fmt.Scanf("%d", &angka1)
	fmt.Printf("Masukan angka 2: %.0d", angka2)
	fmt.Scanf("%d", &angka2)
	if !(angka1 == angka2) {
		println(false)
	} else {
		println(true)
	}
}
