package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", ":8080")
	var name string

	if err != nil {
		fmt.Println("Error connecting to server", err)
	}

	fmt.Println("Connected to Hydra server")
	defer con.Close()
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	name += ": "

	go func() {
		scanner1 := bufio.NewScanner(con)
		for scanner1.Scan() {
			fmt.Println(scanner1.Text())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Fprintf(con, name+message+"\n")
	}
}
