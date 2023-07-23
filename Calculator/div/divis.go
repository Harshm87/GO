package div

import "fmt"

var result1 float64

func PrintDetails(a, b float64) float64 {
	if b!=0{
		result1 = a / b	
		return result1
		}else{
	fmt.Println("ERROR")
}
return 0
}