// write by Tianle Shu
package main

import (
		"fmt"        //import for input and output
	    "sort"       //import for sort
)

var
(
 	A, B, s int       //Declare A,B,s
)

func mergeSort(){
	fmt.Println("Please input 3 diffrent numbers:") //ask customer input 3 diffrent numbers
	A := make([]int, 3)                         //create a array(length=3)    
	fmt.Scanln(&A[0], &A[1], &A[2])                //input 3 different numbers
	fmt.Println(A)                                  // output array A
	fmt.Println("Please input 3 diffrent numbers:")  //ask customer input 3 diffrent numbers
	B := make([]int,3)                              //create a array(length=3)  
	fmt.Scanln(&B[0], &B[1], &B[2])                    //input 3 different numbers
	fmt.Println(B)                                      // output arrayB                       
	s := append(A,B...)                                  //Join 2 lists
	sort.Ints(s)                                      //sorting slices
	fmt.Println(s)                                    //print sort list

}



func main(){     
	mergeSort()                       //call function mergeSort

}

