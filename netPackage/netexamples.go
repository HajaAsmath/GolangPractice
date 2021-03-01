package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listenr, err := net.Listen("tcp", ":8888")

	if err != nil {
		fmt.Println("Error Connecting port", err)
	}

	for {
		con, err := listenr.Accept()
		fmt.Println("Accepted")
		if err != nil {
			fmt.Println("Error Accepting connection", err)
			break
		}
		go func(c net.Conn) {
			defer c.Close()
			clientreader := bufio.NewScanner(con)
			for clientreader.Scan() {
				fmt.Println(clientreader.Text())
			}

		}(con)
	}
}
