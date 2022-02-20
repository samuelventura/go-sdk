package main

import (
	"fmt"
)

type Address struct {
	Line1      string
	Line2      string
	PostalCode uint
	Country    string
}

func main() {
	name := "Jaz"
	fmt.Printf("Hello %s\n", name)
	name = "Jael"
	fmt.Printf("Hello %s\n", name)
	fmt.Println("Hello world!", name)
	count := 1099
	fmt.Printf("Hello %06X %s\n", count, name)

	array := []byte{255, 2, 3, 4}
	fmt.Println("array", array)

	map1 := make(map[string]int)
	map1["one"] = 1

	addr := &Address{}
	addr.Line1 = "Hidalgo 14"
	addr.Line2 = "Col. Centro"
	addr.PostalCode = 78000
	addr.Country = "MX"
}
