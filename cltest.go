package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
	change the 'main' into 'main' and type in console
	a) go build cltest.go
	b) .\cltest 24 13 45 67
	The response will be app executable and the values 
	[C:\xampp\htdocs\go\free_code_camp_beginner\cltest.exe 24 13 29]
	Max: 67
	User's console input: [24 13 45 67]
*/
func mains(){
	
	fmt.Println(os.Args)
	args := os.Args[1:]
	var iArgs = []int{}
	max := 0
	
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	for _, val := range iArgs {
		if val >= max {
			max = val
		}
	}

	fmt.Println("Max:", max)
	fmt.Println("User's console input:", args)
}