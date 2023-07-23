package main

import "fmt"

func main() {
	var name1 string
	fmt.Println("Enter Your name")
	fmt.Scanln(&name1)

	fmt.Println("Hello ",name1)
}