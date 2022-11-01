package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GuessingGame() {
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
}