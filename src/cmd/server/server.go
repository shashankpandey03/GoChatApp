package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("--- Starting Go Chat Server ---")

	listener, err := net.Listen("tcp", ":9960")
	if err != nil {
		fmt.Println("Error starting Go Chat server at post 9960", err)
		return
	} else {
		//channel := make(chan []byte)
		for {
			connection, err := listener.Accept()
			if err != nil {
				fmt.Println("Error while accepting connections", err)
				return
			}

			fmt.Println("Accepting connection")

			go handleConnection(connection)
		}
	}
}

func handleConnection(connection net.Conn) {
	addr := connection.RemoteAddr()
	fmt.Println("*** Received message from", addr.String())

	message := make([]byte, 1024)
	_, err := connection.Read(message)
	if err != nil {
		fmt.Println("Error while reading message from connection", err)
		//channel <- []byte("failure")
	} else {
		fmt.Println("Message received :", string(message))
		//channel <- []byte("success")
	}
	connection.Close()
	//fmt.Println(<-channel)
}
