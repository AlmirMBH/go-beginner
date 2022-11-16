package main // package import that says that the code will be run in the terminal

import (
	"fmt" 
)

var printLine = fmt.Println

func main() {

    printLine("\nOOPSHIT")
    mainOOP()

	printLine("\nCOLLECT USER INPUT VIA CONSOLE")
    collectInput()

	printLine("\nCASTING")
    casting()

    printLine("\nIF STATEMENT")
    ifStatement()

    printLine("\nSTRINGS")
    stringFunctions()

    printLine("\nRUNES")
    runes()

    printLine("\nTIME")
    timeFunctions()

    printLine("\nOPERATORS")
    operators()

    printLine("\nMATH FUNCTIONS")
    mathFunctions()

    printLine("\nFORMATTED PRINT")
    formattedPrint()

    printLine("\nFOR AND WHILE")
    forWhile()

    printLine("\nARRAY")
    array()

    printLine("\nSLICE ARRAY")
    sliceArray()

    printLine("\nINTEGERS AND POINTERS")
    IntegersPointers()

    printLine("\nARRAYS AND POINTERS")
    arrayPointers()

    printLine("\nFUNCTIONS")
    sum := getSum(4, 5)
    printLine("The function sum is:", sum)
    printLine(getTwoValues(4))
    printLine(getQuotient(3.44, 5.22))
    printLine("The variadic function sum is:", getVariadicSum(5, 3, 4, 5))

    varArr := []int{1, 2, 3, 4, 5, 6}
    printLine("The array function sum is:", getArraySum(varArr)) 

    printLine("\nFILES")
    files()
    fileAppend()

    printLine("\nCLTEST")
    mains()

    printLine("\nMAPS")
    maps()

    printLine("\nGUESSING GAME")
    guessingGame()

}


// 01:32:22