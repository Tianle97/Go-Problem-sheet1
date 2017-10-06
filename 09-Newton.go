//write by Tianle Shu
package main

import (
	"fmt"
	"math"
)

func newton(x float64, z float64)float64{   //the function test the number
	
	return ( z - ((z*z - x) / (2 * z)))   //newton formula
}

func main(){

	var a float64   //declare a 
	fmt.Println("Newton Theory Square Root Calculate") 
	fmt.Println("Please input a number: ")   //ask customer input a number
	fmt.Scanln(&a)                            //input a number
	z := float64(1)
	for z = 1.0; z != newton(a,z); z = newton(a,z){
	      fmt.Printf("The square of ", a)
	fmt.Printf("is ", z)
	fmt.Println(" according to Newtons Theory")  //output newton result
	}
	
	fmt.Printf("The square of", a)
	fmt.Printf("is ", math.Sqrt(a))
	fmt.Println(" according to math.Sqrt")    //output math.Sqrt result

}