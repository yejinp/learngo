package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<- done
	fmt.Println("done")
}
