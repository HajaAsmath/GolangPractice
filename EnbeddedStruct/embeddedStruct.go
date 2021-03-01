package main

import (
	"fmt"
	"math/rand"
	"time"
)

type customRand struct {
	*rand.Rand
	count int
}

func NewCustomRaand(i int64) *customRand {
	return &customRand{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (customRand *customRand) NewRange(min, max int) int {
	customRand.count++
	return customRand.Rand.Intn(max-min) + min
}

func (customRand *customRand) GetCount() int {
	return customRand.count
}

func main() {
	cust := NewCustomRaand(time.Now().UnixNano())
	num := cust.NewRange(10, 20)
	fmt.Println(num)
	fmt.Println(cust.GetCount())
}
