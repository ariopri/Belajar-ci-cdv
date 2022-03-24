package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	forDua()

}

func forDua() {
	person := make(map[string]string)
	person["name1"] = "sukuna"
	person["name2"] = "Yuji"
	person["name3"] = "Gojo"
	for index, value := range person {
		fmt.Println(index, value)
	}
}
