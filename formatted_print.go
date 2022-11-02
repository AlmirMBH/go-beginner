package main

import(
	"fmt"
)

func formattedPrint(){
	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.14786786)
	fmt.Printf("%9.f\n", 3.14224444454)
}