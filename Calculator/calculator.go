package main

import (
	"fmt"
	"thisisatest/add"
	"thisisatest/div"
	"thisisatest/multi"
	"thisisatest/sub"
)

func main(){
	var a, b float64
	var oper string
	var result float64 = 0.0

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&oper)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b)

	switch oper{
	case "+":
		result = add.PrintDetails(a,b)
		fmt.Println(a, oper ,b ,"=",result)
	case "/":
		result = div.PrintDetails(a,b)
		fmt.Println(a, oper ,b ,"=",result)
	case "*":
		result = multi.PrintDetails(a,b)
		fmt.Println(a, oper ,b ,"=",result)
	case "-":
		result = sub.PrintDetails(a,b)
		fmt.Println(a, oper ,b ,"=",result)
	}
	

}