// Write by Tianle Shu
package main

import "fmt"

func main(){
	//var i int 
	var n int
	var l int = 1
	println("Please input a number(1~100): ")       
	fmt.Scanln(&n)											//input a number(1~100)
	fmt.Print("you guess ",l)
	fmt.Println(" times.")
	for n != 50{                                           //for-loop when n doesn't equal to 50
	if n < 50{											   
		println("It is smaller, try again Please: ")
		fmt.Scanln(&n)
			l = l+1                                        //l is a number that is calculate guess times
			fmt.Print("you guess  ",l)	                    
			fmt.Println(" times.")

			
		}
		
		if n > 50{
			println("It is bigger, try again Please: ")
			fmt.Scanln(&n)
			l = l+1                                          //l is a number that is calculate guess times
			fmt.Print("you guess  ", l)	
			fmt.Println(" times.")
			
		}
	
	    	if n == 50{
			println("It is correct!Thank you for your guess")  
			break;											 //jump out of the loop
		}
	}
			fmt.Println("Thanks!")
		
         

	}

