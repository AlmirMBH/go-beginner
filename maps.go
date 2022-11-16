package main

import(
	"fmt"
)

func maps(){
	// var myMap map [keyType]valueType
	var heroes map [string]string
	heroes = make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"

	superPets := map[int]string{1: "Krypto"}

	villains := map[string]string{
		"Joker":"Arnold Schwarzenagger",
		"Lex":"Lex Luther",
	}
	delete(villains, "Lex")
	for k, v := range villains {
		fmt.Printf("Villain on key %s is %s\n", k, v)
	}

	_, booleanValue := heroes["Batman"] 
	printLine("Batman exists:", booleanValue)
	printLine(heroes["Batman"])
	printLine(villains["Joker"], "\n"+villains["Lex"])
	printLine(superPets[1])
	
}