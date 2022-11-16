package main

import(
	"fmt"
	"reflect"
	"strconv"
)

func casting(){
	var1 := 1.5
	var2 := int(var1)
	printLine("Float casted to integer", var2)

	var3 := "500000"
	var4, err := strconv.Atoi(var3)

	printLine("String casted to integer", var4, err, reflect.TypeOf(var4))

	var5 := 500000
	var6 := strconv.Itoa(var5)
	printLine("Integer casted to string", var6, reflect.TypeOf(var6))

	var7 := "5.76"
	var8, err := strconv.ParseFloat(var7, 64)

	if err == nil {
		printLine("String casted to float", var8, reflect.TypeOf(var8))
	}

	if var8, err := strconv.ParseFloat(var7, 64); err == nil {
		printLine("Operation and condition can be written in If statement. The float is:", var8)
	}

	var9 := fmt.Sprintf("%f", 3.14)
	printLine("Float casted to string", var9, reflect.TypeOf(var9))
}