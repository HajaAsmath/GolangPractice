package main

import (
	"fmt"
	"sync"
)

func add(wg sync.WaitGroup, n *int) {
	defer wg.Done()
	*n = *n + 1
}

type data struct {
	i int
	j int
	sync.Mutex
}

func (sc *data) increment(w *sync.WaitGroup) {
	//sc.j = sc.j + 1
	defer w.Done()
	sc.Lock()
	sc.i++
	sc.Unlock()
}

func (sc *data) decrement(w *sync.WaitGroup) {
	//sc.j = sc.j - 1
	defer w.Done()
	sc.Lock()
	sc.i--
	sc.Unlock()
}

func (sc *data) getValue() int {
	//k := sc.j
	sc.Lock()
	v := sc.i
	sc.Unlock()
	return v
}

func main() {
	var wg sync.WaitGroup
	// var sc sync.Mutex
	// var n int
	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		sc.Lock()
	// 		n = n + 1
	// 		sc.Unlock()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(n)

	s := new(data)
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go s.increment(&wg)
		go s.decrement(&wg)
	}
	wg.Wait()
	fmt.Println(s.getValue())
}
