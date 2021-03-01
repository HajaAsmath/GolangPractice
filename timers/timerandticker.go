package main

import (
	"fmt"
	"sync"
	"time"
)

func handle(stop chan bool, c chan bool) {
loop:
	for {
		select {
		case <-stop:
			fmt.Println("in stop")
			break loop
		case <-c:
			time.Sleep(5 * time.Second)
		}
		fmt.Println("exited from select")
	}
	fmt.Println("exited from loop")
}

func tickerExp(ticker *time.Ticker, done chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case <-done:
			fmt.Println("Done")
			wg.Done()
			return
		case ct := <-ticker.C:
			fmt.Println("time is", ct.String())
		}
	}
}

func main() {
	// c := make(chan bool)
	// stop := make(chan bool)
	// go handle(stop, c)
	// for i := 0; i < 10; i++ {
	// 	c <- true
	// 	fmt.Println(i)
	// }
	// stop <- true
	// time.Sleep(10 * time.Second)
	h, _ := time.ParseDuration("2h30m")
	var wg sync.WaitGroup
	fmt.Println(h.Hours())
	fmt.Println(h.String())
	t := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	defer t.Stop()
	wg.Add(1)
	go tickerExp(t, done, &wg)
	time.Sleep(10 * time.Second)
	t.Reset(time.Second * 2)
	time.Sleep(time.Second * 10)
	done <- true
	wg.Wait()
}
