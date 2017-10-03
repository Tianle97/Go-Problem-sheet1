package main

import (
	"fmt"
)

func main (){
	
    var s int = 0
	var n int = 1
	var r int = 0
	for  i := 0; i <=100; i++{     //loop for 0<=i<=100
		    i = i+1
			n = n*i                // we can get n=100!
			}
			for n != 0{             //a second 'loop for' for sum of the digits 
				s = n%10          //s is the remainder of n divisible by 10
				r = r+s			  //r is the sum os	
				n = n/10            // n divisible by 10 takes integer
			}
				  
	fmt.Println(r)
}

