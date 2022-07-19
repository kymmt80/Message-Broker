package main

import (
	"fmt"
	"strconv"
	"sync"
)

type MultipleClientBroker struct{
	messageQueue []chan string
}

func NewMultipleClientBroker() *MultipleClientBroker {
	return &MultipleClientBroker{[]chan string{make(chan string,4),make(chan string,4),make(chan string,4)}}
}

func (b* MultipleClientBroker)send(message string, clientId int){
	if clientId<3 {
		b.messageQueue[clientId]<-message
	}
}

func (b* MultipleClientBroker)recieve(clientId int) string{
	if clientId<3 {
		return <-b.messageQueue[clientId]
	}
	return "error"
}


var wg sync.WaitGroup

func server(b* MultipleClientBroker){
	for i := 0; i < 10; i++{
		fmt.Printf("Server is sending: test"+strconv.Itoa(i)+" to Client "+strconv.Itoa(i%3)+"\n")
		b.send("test"+strconv.Itoa(i)+"\n",i%3); 
	}
	wg.Done()
}

func client(b* MultipleClientBroker,num int){
	for i := 0; i < 3; i++{
		fmt.Printf("Client "+strconv.Itoa(num) + " Recieve: "+b.recieve(num))
	}
	wg.Done()
}

func main() {

	b := NewMultipleClientBroker()

	wg.Add(4)
	go server(b)
	go client(b,0)
	go client(b,1)
	go client(b,2)

	wg.Wait()

}


