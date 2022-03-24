package main

import "fmt"

func main() {
	var name1, name2 string
	fmt.Printf("Masukan nama 1 %s", name1)
	fmt.Scanf("%s", &name1)
	fmt.Printf("Masukan nama 2 %s", name2)
	fmt.Scanf("%s", &name2)

	var result bool = name1 == name2
	fmt.Println(result)

}
