package main

import (
	Countdown "main/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &Countdown.ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown.Countdown(os.Stdout, sleeper)
}
