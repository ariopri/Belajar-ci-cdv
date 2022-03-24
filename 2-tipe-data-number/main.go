package main

import "fmt"

func main() {
	newFunction()

}

func newFunction() {
	var (
		a      int
		b      int
		c      int
		result int
	)
	newFunction1(a, b, c, result)
}

func newFunction1(a int, b int, c int, result int) {
	newFunction3(a, b, c)
	newFunction2(result, a, b, c)
}

func newFunction3(a int, b int, c int) {
	newFunction4(a, b, c)
}

func newFunction4(a int, b int, c int) {
	newFunction7(a)
	newFunction6(b)
	newFunction5(c)
}

func newFunction7(a int) {
	fmt.Printf("masukan nilai a : %.0d", a)
	fmt.Scanf("%d", &a)
}

func newFunction6(b int) {
	fmt.Printf("\n masukan nilai b:%.0d", b)
	fmt.Scanf("%d", &b)
}

func newFunction5(c int) {
	fmt.Printf("masukan nilai c: %.0d", c)
	fmt.Scanf("%d", &c)
}

func newFunction2(result int, a int, b int, c int) {
	result = a + b + c
	fmt.Println(result)
}
