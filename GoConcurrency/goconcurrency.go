package main

import (
	"fmt"
	"math/rand"
	"time"
)

var securitylevel = map[string]int{
	"James": 10,
	"Rio":   20,
	"Mat":   20,
}

func main() {
	c := make(chan bool, 2)
	//ch := make(chan bool, 2)
	str := "Hello"
	go waitAndSay(c, str, 2)
	if b := <-c; b {
		fmt.Println("World")
	}
	if bb := <-c; bb {
		fmt.Println("World")
	}
	fmt.Println("Done")
	rand.Seed(time.Now().UnixNano())

	c1 := make(chan int)
	c2 := make(chan int)

	name := "James"

	go getsecuritylevel(c1, name)
	go getsecuritylevel(c2, name)

	select {
	case s1 := <-c1:
		fmt.Println(name, "has security level", int(s1), "found in server 1")
	case s2 := <-c2:
		fmt.Println(name, "has security level", int(s2), "found in server 2")
	case <-time.After(time.Second * 11):
		fmt.Println("time out")
	}

}

func getsecuritylevel(c chan int, name string) {
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	c <- securitylevel[name]
}

func waitAndSay(c chan bool, str string, num int) {
	for n := 0; n < num; n++ {
		c <- true
		fmt.Println(str)
	}
}
