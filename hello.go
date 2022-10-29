package main // package import that says that the code will be run in the terminal

import (
	"bufio"
	"fmt" // format
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var printLine = fmt.Println // instead of using fmt.Println() we can use an alias

func main() {

	// collect user input via console
	printLine("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		printLine("Hello", name)
		printLine("The type of the enter value is:", reflect.TypeOf(name))
		printLine("The type of the additional value is:", reflect.TypeOf('@'))
	} else {
		log.Fatal(err)
		// printLine(err)
	}

	// casting
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

	// if statement
	iAge := 33
	if (iAge >= 1) && (iAge <= 18) {
		printLine("Within age")
	} else if iAge < 30 {
		printLine("Older than 22")
	} else {
		printLine("Within specified age:", !true)
	}

	// strings
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

	// runes
	rStr := "abcdefg"
	printLine("Rune count", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c \n", i, runeVal, runeVal)
	}

	// time
	now := time.Now()
	printLine("Raw now:", now)
	printLine("Unix:", now.Unix())
	printLine("Date now parsed:", now.Year(), now.Month(), now.Day())
	printLine("Time now parsed:", now.Hour(), now.Minute(), now.Second())
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	printLine("Random:", randNum)

	// operators
	printLine("5+4", 5+4)
	mInt := 1
	mInt = mInt + 1
	mInt += 1
	mInt++
	printLine(mInt)

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
}
