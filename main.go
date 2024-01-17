package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			message <- i
		}
		close(message)
	}()

	time.Sleep(3 * time.Second)

	for i := 0; i < 100; i++ {
		msg := <-message
		fmt.Println(msg)
	}
}
