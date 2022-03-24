package main

import "fmt"

func main() {
	var person = []map[string]string{
		{"nama": "madara", "clan": "uciha"},
		{"nama": "hashirama", "clan": "senju"},
	}
	for _, value := range person {
		fmt.Println(value["nama"], value["clan"])
	}

}
