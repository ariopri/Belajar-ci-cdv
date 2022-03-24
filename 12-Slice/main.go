package main

import "fmt"

func main() {
	// bikin slice sederhana
	slice := []string{
		"madara", "Hashirama Senju",
	}
	fmt.Println(slice[0], "\n", slice[1])

	Slice()

}

func Slice() {
	names := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	slice1 := names[3:6]
	fmt.Println(slice1)
	names[3] = "hari kamis"
	fmt.Println(names[3])
	// bikin slice baru
	newSlice := make([]string, 2, 3)
	newSlice[0] = "madara"
	newSlice[1] = "Hashirama"
	fmt.Println(newSlice)
	fmt.Println(newSlice[0])
	fmt.Println(newSlice[1])
	// menambahkan Slice
	newSlice2 := append(newSlice, "Jujutsu Kaisen")
	// newSlice2[0] = "yuji"
	fmt.Println(newSlice2)
}
