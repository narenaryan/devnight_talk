package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func compress(file string) {
	fmt.Println("Compressing: ", file)
	time.Sleep(2 * time.Second)
}

func main() {
	for _, file := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}

	fmt.Println("See main is not blocked")

	wg.Wait()
}
