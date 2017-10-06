// Write by Tianle Shu
package main

import "fmt"
import "strings"
	
func Palindrome(s string)string{   //the function to test if the word is a palindrome
	var m int                    // declare int m that is the middle position of the word
	var l int                    //declare int l 
	m = len(s) / 2                // m is middle of the length of s
	l = len(s) -1                // l is the last of the length of s
	for i := 0; i < m; i++{          
		if s[i] != s[l-i]{               //check if in the middle of each letter are the same
			return("have not palindrome.")
		}
	}
	return("have a palindrome. ")
	
}

func main(){

		var p string
		fmt.Println("Please input a string: ")     //ask customer input  word
		fmt.Scanln(&p)                              //input word
		p = strings.ToLower(p)                      //convert to lowercase
		fmt.Println(Palindrome(p))                  //call function
	}