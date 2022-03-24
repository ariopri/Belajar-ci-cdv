package main

import "fmt"

func main() {
	slice := []string{"a", "r", "i", "o"}

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])

	// menambahkan data
	slice = append(slice, "n")
	fmt.Println(slice)
	slice = append(slice, "i", "t")
	fmt.Println(slice)

	// menghapus data
	slice = append(slice[:2], slice[4:]...)
	fmt.Println(slice)

	// mengubah data
	slice[1] = "r"
	fmt.Println(slice)

}
