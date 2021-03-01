package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	con, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("Error in dialing", err)
	}
	defer con.Close()
	writer := bufio.NewWriter(con)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		writer.WriteString(str + "\n")
		writer.Flush()
	}
}
