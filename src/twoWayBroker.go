package main

import (
	"fmt"
	"strconv"
)

type TwoWayBroker struct{
	cmessageQueue chan string
	smessageQueue chan string
}

func NewTwoWayBroker() *TwoWayBroker {
	return &TwoWayBroker{make(chan string,4),make(chan string,4)}
}

func (b* TwoWayBroker)ssend(message string){
	b.cmessageQueue<-message
}

func (b* TwoWayBroker)srecieve() string{
	return <-b.smessageQueue
}

func (b* TwoWayBroker)csend(message string){
	b.smessageQueue<-message
}

func (b* TwoWayBroker)crecieve() string{
	return <-b.cmessageQueue
}


func server(b* TwoWayBroker){
	for i := 0; i < 4; i++{
		fmt.Printf("Server Recieve: "+b.srecieve())
	}
	for i := 0; i < 4; i++{
		fmt.Printf("Server is sending: test"+strconv.Itoa(i)+"\n")
		b.ssend("test"+strconv.Itoa(i)+"\n"); 
	}
}

func main() {

	b := NewTwoWayBroker()

	go server(b)

	for i := 0; i < 4; i++{
		fmt.Printf("Client is sending: "+strconv.Itoa(i)+"\n")
		b.csend(strconv.Itoa(i)+"\n")
	}
	//Client Code
	for i := 0; i < 4; i++{
		fmt.Printf("Client Recieve: "+b.crecieve())
	}

}


