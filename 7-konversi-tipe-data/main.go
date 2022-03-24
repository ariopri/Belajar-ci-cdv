package main

import "fmt"

func main() {
	var nilai32 int32 = 3222
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai64)
	fmt.Println(nilai8)
	//ini kalau int
	var name string = "madara"
	// fmt.Printf("Masukan namanya disini: %s", name)
	// fmt.Scanf("%s", name)
	fmt.Println(name)

	var e = name[0]
	var eString = string(e)
	fmt.Println(eString)
}
