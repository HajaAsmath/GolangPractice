package HydraChat

import (
	"bufio"
	"net"
)

type client struct {
	*bufio.Reader
	*bufio.Writer
	wc chan string
}

//Starting a client once request is received
func StartClient(messageChannel chan string, con net.Conn, done chan bool) *client {
	client := &client{
		Reader: bufio.NewReader(con),
		Writer: bufio.NewWriter(con),
		wc:     make(chan string),
	}

	go func() {
		scanner := bufio.NewScanner(client.Reader)
		for scanner.Scan() {
			logger.Println("Message" + scanner.Text())
			messageChannel <- scanner.Text()
		}
		done <- true
	}()
	go client.writeMonitor()
	return client
}

func (client *client) writeMonitor() {
	for message := range client.wc {
		client.WriteString(message + "\n")
		client.Flush()
	}
}
