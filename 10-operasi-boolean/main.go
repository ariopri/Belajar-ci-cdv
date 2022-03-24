package main

import "fmt"

func main() {
	var nilai int
	fmt.Printf("Masukan nilai %.0d", nilai)
	fmt.Scanf("%.0d", nilai)
	var (
		lulusNilaiAkhir bool = nilai >= 80
		lulus           bool = lulusNilaiAkhir
	)
	fmt.Println(lulus)
}
