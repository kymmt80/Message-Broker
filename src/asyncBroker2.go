package main

import (
	"fmt"
	"strconv"
	"time"
)


type AsyncBroker struct{
	messageQueue chan string
}

func NewAsyncBroker() *AsyncBroker {
	return &AsyncBroker{make(chan string,4)}
}

func (b* AsyncBroker)send(message string){
	b.messageQueue<-message
}

func (b* AsyncBroker)recieve() string{

	return <-b.messageQueue
}


func server(b* AsyncBroker){
	for i := 0; i < 10; i++{
		fmt.Printf("Server is sending: test"+strconv.Itoa(i)+"\n")
		b.send("test"+strconv.Itoa(i)+"\n"); 
	}
}

func main() {

	b := NewAsyncBroker()

	//b=NewAsyncBroker()
	go server(b)


	time.Sleep(100000000)
	fmt.Printf("client woke up\n")

	for i := 0; i < 10; i++{
		fmt.Printf("Client Recieve: "+b.recieve())
	}

}