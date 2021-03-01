package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error")
	}
	for {
		conn, _ := l.Accept()
		fmt.Println("Connection Accepted from", conn.RemoteAddr())
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		conn.Write([]byte("Message Received \n"))
	}
}
