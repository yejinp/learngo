package main

import "fmt"

func main() {
	messages := make(chan string)
	
	go func(s string) {messages <- s}("ping pang")

	msg := <-messages
	fmt.Println(msg)
}
