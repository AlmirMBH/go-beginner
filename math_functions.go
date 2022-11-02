package main

import(
	"fmt"
	"math"
)

func mathFunctions(){
	// math
	printLine("Abs(-10)", math.Abs(-10))
	printLine("Pow(4, 2)", math.Pow(4, 2))
	printLine("Sqrt(16)", math.Sqrt(16))
	printLine("Cbrt(64)", math.Cbrt(64))
	printLine("Ceil(4.4)", math.Ceil(4.4))
	printLine("Floor(4.4)", math.Floor(4.4))
	printLine("Round(4.4)", math.Round(4.4))
	printLine("Log2(8)", math.Log2(8))
	printLine("Log10(100)", math.Log10(100))
	// Get the log of e to the power of 2
	printLine("Log(7.389)", math.Log(math.Exp(2)))
	printLine("Max(5, 4)", math.Max(5, 4))
	printLine("Min(5, 4)", math.Min(5, 4))
	// convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 * math.Pi)
	fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)
	printLine("Sin(90) = ", math.Sin(r90)) // radiant version of 90 degrees
}