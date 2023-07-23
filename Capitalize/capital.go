package main

import (
	"fmt"
	"unicode"
)

func TOupper(s string) string{
	a:=[]rune(s)
	for i,c:=range a {
		if unicode.IsLower(c){
			a[i] = unicode.ToUpper(c)
	}
}
return string(a)
}

func main(){
	var input string = "Hello my name is harsh"
	fmt.Println("String Before Capitalization\n ",input)
	var output string = TOupper(input)
	fmt.Println("String after Capitalization:\n ",output)

}