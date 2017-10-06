//write by Tianle Shu
package main

import "fmt"

func reserveString(s string) string{
	var a string                  //declare string a
	for i := len(s)-1; i >= 0; i--{     
		a = a + string(s[i])      //a = reserve of s  
		}
	return a
}


func main(){

	var b string
	fmt.Println("Please input string: ")     //ask customer input string
	fmt.Scanln(&b)                            // input string
	fmt.Println("Reserves String is : ", reserveString(b))  //output reserves String
 
}

