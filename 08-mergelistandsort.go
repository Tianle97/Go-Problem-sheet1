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
	fmt.Println("Please input 3 diffrent numbers:")
	A := make([]int, 3)
	fmt.Scanln(&A[0], &A[1], &A[2])
	fmt.Println(A)
	fmt.Println("Please input 3 diffrent numbers:")
	B := make([]int,3)
	fmt.Scanln(&B[0], &B[1], &B[2])
	fmt.Println(B)
	s := append(A,B...)
	sort.Ints(s)
	fmt.Println(s)

}



func main(){     
	mergeSort()
	
	//fmt.Println("Please input array A :")
	//fmt.Scanln(A...)
	//fmt.Println("%v", A)
	//B := make([]int, 5)           //array B  
	//fmt.Println("Please input array B :")
	//fmt.Scanln(B...)
	//fmt.Println("%v", B)
	//s := append(A, B...)   //append AB
	
 	//sort.Ints(s)           //sort AB in s    
	//fmt.Println(s)
}

