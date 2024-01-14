package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	// time.Sleep(2 * time.Second)
	end := time.Now()

	diff := end.Sub(start)

	fmt.Println(start.String())
	fmt.Print("Result: ")
	fmt.Println(diff)

	future := time.Date(2024, time.March, 28, 12, 12, 12, 12, time.UTC)

	// isAfterFuture := future.After(time.Now())
	fmt.Println(time.ParseDuration(time.Until(future).String()))

}
