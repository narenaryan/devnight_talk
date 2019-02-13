package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int)
	go do(done)
	fmt.Println("I execute without blocking")
	<-done
}

func do(done chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("I did some work concurrently")
	done <- 1
}
