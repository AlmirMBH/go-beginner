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

	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("String to rune loop value %d:\n", v)
	}

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
	// convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 * math.Pi)
	fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)
	printLine("Sin(90) = ", math.Sin(r90)) // radiant version of 90 degrees

	// formatted print
	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.14786786)
	fmt.Printf("%9.f\n", 3.14224444454)

	// for loop
	for x := 1; x < 5; x++ {
		printLine(x)
	}

	// while with for
	fx := 0
	for fx < 5 {
		printLine(fx)
		fx++
	}

	// guessing game
	secs := time.Now().Unix()
	rand.Seed(secs)
	randNumber := rand.Intn(50) + 1

	for true {
		fmt.Print("Guess a number between 0 and 50:")
		// printLine("Random number is: ", randNumber)
		reader := bufio.NewReader(os.Stdin) // standard input
		guess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)

		if err != nil {
			log.Fatal(err)
		}

		if iGuess > randNumber {
			printLine("Pick a Lower Value")
		} else if iGuess < randNumber {
			printLine("Pick a Higher Value")
		} else {
			printLine("Your guess is correct!")
			break
		}
	}

	// arrays
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		printLine("Array element:", num)
	}

	var arr1 [5]int
	arr1[0] = 5
	arr2 := [5]int{1, 2, 3, 4, 5}
	printLine("Index 0:", arr2[0])
	printLine("Array 2 length:", len(arr2))

	for i := 0; i < len(arr2); i++ {
		printLine("Array 2 element in loop", i)
	}

	for i, v := range arr2 {
		fmt.Printf("Value on index %d is %d\n", i, v)
	}

	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("Array 3 element in loop on index i %d and j %d is: ", i, j)
			printLine(arr3[i][j])
		}
	}


    // 1:00:08

}
