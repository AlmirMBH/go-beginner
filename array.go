package main

import(
	"fmt"
)


func array(){
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		printLine("Array element:", num)
	}

	var arr1 [5]int
	arr1[0] = 5
	arr2 := [5]int{1, 2, 3, 4, 5}
	printLine("Index 0:", arr2[0])
	printLine("Array 2 length:", len(arr2))

	for i := 0; i < len(arr2); i++ {
		printLine("Array 2 element in loop", i)
	}

	for i, v := range arr2 {
		fmt.Printf("Value on index %d is %d\n", i, v)
	}


    // matrix
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("Array 3 element in loop on index i %d and j %d is: %d\n", i, j, arr3[i][j])
		}
	}
}