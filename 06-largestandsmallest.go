package main

import "fmt"

func main(){

	var min = 0
	var max = 0
	A := make([]int, 6)
	println("Please input 6 diffrent numbers:")
	
	fmt.Scanln(&A[0], &A[1], &A[2],&A[3], &A[4], &A[5])

	min = A[0]
	max = A[0]

	for i := range A{
		if A[i] < min{
			min = A[i]
		}
		if A[i] > max{
			max = A[i]
		}
	}
	fmt.Println("The max is: ", max)
	fmt.Println("The min is: ", min)

}