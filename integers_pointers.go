package main

import(
	"fmt"
)

func IntegersPointers(){
	// integers & pointers
    f3 := 8
    printLine("F3 before function processing", f3)
    changeValue(&f3)
    printLine("F3 after function processing", f3)

    f4 := 33
    var f4Ptr *int = &f4
    fmt.Printf("F4 TEST %v\n", &f4)
    fmt.Printf("F4 pointer to memory location %v\n", f4Ptr)
    fmt.Printf("F4 pointer value %v\n", *f4Ptr)  

    *f4Ptr = 22
    fmt.Printf("F4 pointer value modified %v\n", *f4Ptr) 
}