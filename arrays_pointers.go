package main

import(
	"fmt"
)

func arrayPointers(){
	 // arrays & pointers
	 pArr := [4]int{1, 2, 3, 4}
	 dblArrVals(&pArr)
	 printLine(pArr)
 
	 isSlice := []float64{11, 13, 17}
	 fmt.Printf("Average: %.3f\n", getAverage(isSlice ...))
}