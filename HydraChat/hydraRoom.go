package HydraChat

import (
	"net"
	"sync"
)

type room struct {
	name           string
	messageChannel chan string
	*sync.RWMutex
	client map[chan<- string]struct{}
	quit   chan bool
}

//Configure a chat room server where clients can join
func CreateRoom(name string) *room {
	room := &room{
		name:           name,
		messageChannel: make(chan string),
		RWMutex:        new(sync.RWMutex),
		client:         make(map[chan<- string]struct{}),
		quit:           make(chan bool),
	}
	room.Run()
	return room

}

func (r *room) AddClient(con net.Conn) chan<- string {
	done := make(chan bool)
	clientChan := StartClient(r.messageChannel, con, done)
	r.client[clientChan.wc] = struct{}{}
	go func() {
		<-done
		r.RemoveClient(clientChan.wc)
	}()
	return clientChan.wc
}

func (r *room) RemoveClient(clientChan chan<- string) {
	r.Lock()
	close(clientChan)
	delete(r.client, clientChan)
	r.Unlock()
}

func (r *room) clientCount() int {
	return len(r.client)
}

func (r *room) Run() {
	go r.broadcastMessage()
}

func (r *room) broadcastMessage() {
	r.RLock()
	defer r.RUnlock()
	for message := range r.messageChannel {
		for wc, _ := range r.client {
			wc <- message
		}
	}
}
