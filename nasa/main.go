package main

import (
	"fmt"
	"time"
)

func main() {
	isAbort := make(chan bool)

	go runAbortScreen(isAbort)

	select {
	case <-isAbort:
		return
	case <-time.After(5 * time.Second):
		// Come out of select
	}

	launch()
}

func launch() {
	fmt.Println("AOK! Rocket is launched! God save us")
}

func runAbortScreen(isAbort chan bool) {
	fmt.Println("Hit any key to abort launch")
	var key rune
	fmt.Scanf("%s", &key)
	isAbort <- true
}
