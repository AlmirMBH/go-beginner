package main

import(
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func files() {
    f, err := os.Create("files/data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    iPrimeArr := []int{2, 3, 5, 7, 11}
    var sPrimeArr []string
    for _, i := range iPrimeArr {
        sPrimeArr = append(sPrimeArr, strconv.Itoa(i)) // int -> str
    }
    for _, num := range sPrimeArr {
        _, err := f.WriteString(num + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }

    f, err = os.Open("files/data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scan1 := bufio.NewScanner(f)
    for scan1.Scan() {
        printLine("Prime: ", scan1.Text())
    }

    if err := scan1.Err(); err != nil {
        log.Fatal(err)
    } 
}

func fileAppend(){
	_, err := os.Stat("files/data.txt") // Stat resolves the path relative to the root of your project directory
    if errors.Is(err, os.ErrNotExist) {
        printLine("File does not exist")
    }else{
		f, err := os.OpenFile("files/data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0777) // 0644 - access permissions
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}