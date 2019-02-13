package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)
	go Player("ping", table)
	go Player("pong", table)

	ball := new(Ball)
	table <- ball

	<-time.After(5 * time.Millisecond)

}

func Player(message string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(message, ball.hits)
		table <- ball
	}
}
