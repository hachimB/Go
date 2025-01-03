package main

import "fmt"

func str() {
	str := "Hello from 🇳🇪"
	fmt.Println(len(str))
	fmt.Println(string(str[0]))

	for _, char := range str {
		fmt.Printf("%c", char)
	}
	fmt.Printf("\n")

	var s string
	for i := 0; i < len(str); i++ {
		s += string(str[i])
	}
	fmt.Println(s)
	fmt.Printf("\n")
}
