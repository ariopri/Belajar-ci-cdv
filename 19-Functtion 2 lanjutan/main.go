package main

import "fmt"

func main() {
	//function valuee
	/*
		function adalah first class citizen
		function juga merupakan tipe data dan bisa disimpan di variabel
	*/
	functionvalue := ContohFunctionValue
	fmt.Println(functionvalue("Itadori Yuji"))
	fmt.Println(ContohFunctionValue("Gojo Satoru"))
}

func ContohFunctionValue(name string) string {
	/*
		nah ContohFunctionValue bisa kita jadikan / dibuat sebagai variabel
		dan bisa kita panggil dengan menggunakan nama variabel tersebut
	*/
	return "Hello " + name
}
