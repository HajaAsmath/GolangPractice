package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	go run()
	go dial()
	time.Sleep(20 * time.Second)
}

func dial() {
	conn, _ := net.Dial("tcp", ":8888")
	fmt.Println(conn.LocalAddr())
	writer := bufio.NewWriter(conn)
	writer.WriteString("Hello" + "\n")
	writer.Flush()

	go func(con net.Conn) {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println("Message from server: ", text)
		}
	}(conn)
}

func run() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Print("Error")
	}
	for {
		conn, _ := l.Accept()
		fmt.Println("Connection Accepted from", conn.RemoteAddr())
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		conn.Write([]byte("Message Received \n"))
	}
}
