# Message Broker
This repository contains basic implementations of message brokers with different characteristics. these message brokers are meant to pass messages between two or more goroutines.

## Source Codes
the source codes are located in ```src/``` folder. bellow is a brief explanation of each file:
- ```asyncBroker.go```: Asynchronous message passing between a client and server.
- ```asyncBroker2.go```: Asynchronous message passing between a client and server with overflow handling.
- ```multipleClientBroker.go```: Asynchronous message passing between several clients and one server.
- ```syncBroker.go```: Synchronous message passing between a client and server.
- ```twoWayBroker.go```: Asynchronous two-way message passing between a client and server.