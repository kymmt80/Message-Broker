package main

import (
	"fmt"
	"strconv"
)


type SyncBroker struct{
	messageQueue chan string
	ok chan bool
}

func NewSyncBroker() *SyncBroker {
	return &SyncBroker{make(chan string,4),make(chan bool,1)}
}

func (b* SyncBroker)send(message string){
	b.ok<-false
	b.messageQueue<-message
	for <-b.ok!=true{}
}

func (b* SyncBroker)recieve() string{
	var message =<-b.messageQueue
	b.ok<-true
	return message
}


func server(b* SyncBroker){
	for i := 0; i < 4; i++{
		fmt.Printf("Server is sending: test"+strconv.Itoa(i)+"\n")
		b.send("test"+strconv.Itoa(i)+"\n"); 
	}
}

func main() {

	b := NewSyncBroker()

	go server(b)

	//Client Code:
	for i := 0; i < 4; i++{
		fmt.Printf("Client Recieve: "+b.recieve())
	}

}
