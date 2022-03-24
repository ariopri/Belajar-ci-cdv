package main

import "fmt"

func main() {
	i := newFunction1()
	newFunction(i)

}

func newFunction1() (i int) {
	fmt.Printf("masukan angka: %.0d", i)
	fmt.Scanf("%d", &i)
	return i
}

func newFunction(i int) {
	if i == 10 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
