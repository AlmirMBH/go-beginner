package main

import(
	"math/rand"
	"time"
)

func timeFunctions(){
	now := time.Now()
	printLine("Raw now:", now)
	printLine("Unix:", now.Unix())
	printLine("Date now parsed:", now.Year(), now.Month(), now.Day())
	printLine("Time now parsed:", now.Hour(), now.Minute(), now.Second())
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	printLine("Random:", randNum)
}