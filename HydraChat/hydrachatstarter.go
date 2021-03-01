package HydraChat

import (
	"SingletonPattern/Logger"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger = Logger.GetLoggerInstance()

func Run() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Println("Error creating server", err)
	}
	room := CreateRoom("HydraChat")
	fmt.Println(room)
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		<-ch
		room.quit <- true
		os.Exit(0)
	}()

	for {
		cn, err := listener.Accept()
		if err != nil {
			logger.Println("Error Accepting client")
			break
		} else {
			go handleConnection(room, cn)
		}
	}
}

func handleConnection(room *room, con net.Conn) {
	logger.Println("Adding Client to room")
	room.AddClient(con)
}
