package main

import (
	"fmt"
	"time"
)

func slowCounter(tick *time.Ticker) chan bool {
	done := make(chan bool)
	go func() {
	Loop:
		for {
			select {
			case time := <-tick.C:
				fmt.Println(time)
			case <-done:
				break Loop
			}
		}
	}()
	return done
}
func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := slowCounter(ticker)
	time.Sleep(10 * time.Second)
	done <- true
}
