package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math"
	"strconv"
)
	


// functions
func GetAverage(nums ...float64) float64 {
    var sum float64 = 0.0
    var numSize float64 = float64(len(nums))

    for _, val := range nums{
        sum += val
    }
    return sum / numSize
}


func GetSum(x int, y int) int {
    return x + y
}


func GetTwoValues(x int) (int, int) {
    return x + 1, x + 2
}


func GetQuotient(x float64, y float64) (ans float64, err error) {
    if y == 0 {
        return 0, fmt.Errorf("you cannot divide by zero")
    }else{
        return math.Round(x / y), nil
    }
}

// variadic functions (receiving an unknown number of values)
func GetVariadicSum(nums... int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}


func GetArraySum(arr []int) int {
    sum := 0
    for _, num := range arr {
        sum += num
    }
    return sum
}


// pointers
func ChangeValue(f3 *int) int {
    *f3 = 12
    return *f3
}

func DblArrVals(arr *[4]int) {
    for x := 0; x < len(arr); x++ {
        arr[x] *= 2
    }
}


// files
func CreateFile() {
    f, err := os.Create("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    iPrimeArr := []int{2, 3, 5, 7, 11}
    var sPrimeArr []string
    for _, i := range iPrimeArr {
        sPrimeArr = append(sPrimeArr, strconv.Itoa(i)) // int -> str
    }
    for _, num := range sPrimeArr {
        _, err := f.WriteString(num + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }

    f, err = os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scan1 := bufio.NewScanner(f)
    for scan1.Scan() {
        printLine("Prime: ", scan1.Text())
    }

    if err := scan1.Err(); err != nil {
        log.Fatal(err)
    }
}


