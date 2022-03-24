package main

import "fmt"

func main() {
	belajarSwitch()
}

func belajarSwitch() {
	var name string
	fmt.Printf("Masukan namanya disini: %s", name)
	fmt.Scanf("%s", &name)

	switch name {
	case "Ario":
		fmt.Printf("Selamat datang bos %s", name)
		break
	case "Gojo":
		fmt.Printf("Guru Itadori Yuji %s", name)
		break
	case "yuji":
		fmt.Printf("Mc anime jujutsu kaisen adalah %s", name)
		break
	default:
		fmt.Printf("Nama %s tidak terdaftar", name)
	}

}
