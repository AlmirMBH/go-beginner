package main

import(
	"fmt"
	"strings"
)

func stringFunctions(){
	sV1 := " A w o rd "
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	printLine(sV2)
	printLine("Length:", len(sV2))
	printLine("Contains Another:", strings.Contains(sV2, "other"))
	printLine("O at index:", strings.Index(sV2, "o"))
	printLine("Replace 'o'", strings.Replace(sV2, "o", "T", -1))
	printLine("Trim white spaces:", strings.TrimSpace(sV2))
	printLine("Split by '-':", strings.Split("a-b-c-d", "-"))
	printLine("String to lower", strings.ToLower(sV2))
	printLine("String to lower", strings.ToUpper(sV2))
	printLine("Prefix", strings.HasPrefix("Almir", "Alm"))
	printLine("Prefix", strings.HasSuffix("Almir", "mir"))

	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("String to rune loop value %d:\n", v)
	}
}