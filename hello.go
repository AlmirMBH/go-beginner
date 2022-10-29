package main // package import that says that the code will be run in the terminal

// import "fmt"

import(
    "fmt" // format
    "bufio"
    "log"
    "os"    
    "reflect"    
    "strconv"
    "strings"
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

    if  err == nil{
        printLine("String casted to float", var8, reflect.TypeOf(var8))
    }

    if  var8, err := strconv.ParseFloat(var7, 64); err == nil{
        printLine("Operation and condition can be written in If statement. The float is:", var8)
    }

    var9 := fmt.Sprintf("%f", 3.14)
    printLine("Float casted to string", var9, reflect.TypeOf(var9))


    // if statement
    iAge := 33
    if (iAge >= 1) && (iAge <= 18){
        printLine("Within age")
    } else if iAge < 30 {
        printLine("Older than 22")
    } else{
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
    
}

