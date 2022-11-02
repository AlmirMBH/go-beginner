package main

import (
	"fmt"
	"math"
)
	


// functions
func getAverage(nums ...float64) float64 {
    var sum float64 = 0.0
    var numSize float64 = float64(len(nums))

    for _, val := range nums{
        sum += val
    }
    return sum / numSize
}


func getSum(x int, y int) int {
    return x + y
}


func getTwoValues(x int) (int, int) {
    return x + 1, x + 2
}


func getQuotient(x float64, y float64) (ans float64, err error) {
    if y == 0 {
        return 0, fmt.Errorf("you cannot divide by zero")
    }else{
        return math.Round(x / y), nil
    }
}

// variadic functions (receiving an unknown number of values)
func getVariadicSum(nums... int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}


func getArraySum(arr []int) int {
    sum := 0
    for _, num := range arr {
        sum += num
    }
    return sum
}


// pointers
func changeValue(f3 *int) int {
    *f3 = 12
    return *f3
}

func dblArrVals(arr *[4]int) {
    for x := 0; x < len(arr); x++ {
        arr[x] *= 2
    }
}


