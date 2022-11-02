package main

func forWhile(){
	for x := 1; x < 5; x++ {
		printLine(x)
	}

	// while with for
	fx := 0
	for fx < 5 {
		printLine(fx)
		fx++
	}
}