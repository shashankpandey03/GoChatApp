package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("--- Starting Go Chat Client ---")

	var conn net.Conn

	for {
		reader := bufio.NewReader(os.Stdin)

		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading message", err)
			panic(err)
		} else {
			if message != "quit" {

				conn, err = net.Dial("tcp", "localhost:9960")
				if err != nil {
					fmt.Println("Error while conencting to server", err)
					panic(err)
				}

				fmt.Println("Sending message to server :", message)

				_, err = conn.Write([]byte(message))
				if err != nil {
					fmt.Println("Error while writing message to server", err)
					panic(err)
				} else {
					fmt.Println("Message sent to server")
					conn.Close()
				}
			} else {
				conn.Close()
			}
		}
	}
}
