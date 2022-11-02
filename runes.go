package main

import(
	"fmt"
	"unicode/utf8"
)


func runes(){
	rStr := "abcdefg"
	printLine("Rune count", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c \n", i, runeVal, runeVal)
	}
}