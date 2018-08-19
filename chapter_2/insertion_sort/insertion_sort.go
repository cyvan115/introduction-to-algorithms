package main

import (
	"fmt"
)

func main() {
	A := []int{5, 4, 3, 2, 1}
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
	fmt.Println(A)
}
