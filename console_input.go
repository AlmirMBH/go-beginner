package main

import(
	"bufio"
	"log"
	"os"
	"reflect"	
)

func collectInput(){
	printLine("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		printLine("Hello", name)
		printLine("The type of the entered value is:", reflect.TypeOf(name))
		printLine("The type of the additional value is:", reflect.TypeOf('@'))
	} else {
		log.Fatal(err)
		// printLine(err)
	}
}