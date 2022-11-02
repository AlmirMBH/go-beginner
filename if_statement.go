package main


func ifStatement(){
	iAge := 33
	if (iAge >= 1) && (iAge <= 18) {
		printLine("Within age")
	} else if iAge < 30 {
		printLine("Older than 22")
	} else {
		printLine("Within specified age:", !true)
	}
}