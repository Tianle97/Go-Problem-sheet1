package main

import "fmt"

func read(n int) ([]int, error) {
	in := make([]int, n)
	for i := range n {
	  _, err := fmt.Scan(&in[i])
	  if err != nil {
		 return in[:i], err
	  }
	}
	return in, nil
  }