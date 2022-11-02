package main

import(
	"fmt"
	"reflect"
)

func sliceArray(){
	// slice array
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:]) // [:] take all, [:2] slice and take the first two elements
	printLine("This is a byte array converted to", reflect.TypeOf(bStr), bStr)

    sl1 := make([]string, 5)
    sl1[0] = "Society"
    sl1[1] = "of"
    sl1[2] = "the"
    sl1[3] = "Simulated"
    sl1[4] = "Universe"
    printLine("Array size", len(sl1))

    for i := 0; i < len(sl1); i++ {
        fmt.Printf("SL1 element on index %d is %s\n", i, sl1[i])
    }

    counter := 0
    for _, i := range sl1 {
        fmt.Printf("SL1 element in range on index %d is %s\n", counter, i)
        counter++
    }

    sl2 := []int{1, 2, 3, 4, 5, 6, 7}
    printLine("First three elements", sl2[:3])
    printLine("Last three elements", sl2[3:])
    printLine("Last element", len(sl2)-0)
    sl2 = append(sl2, 12)
    printLine("Append to sl2:", sl2)
}