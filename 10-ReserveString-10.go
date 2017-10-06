package main

import "fmt"

func reserveString(s string) string{
	var a string 
	for i := len(s)-1; i >= 0; i--{
		a = a + string(s[i]) 
		}
	return a
}


func main(){

	var b string
	fmt.Println("Please input string: ")
	fmt.Scanln(&b)
	fmt.Println("Reserves String is : ", reserveString(b))

}

