package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
	listen, err := net.Listen("tcp", ":2021")
	if err != nil {
		log.Fatal("Error", err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				log.Fatal("connect err: ", err)
			}

			clientChan <- rpc.NewClient(conn)
		}
	}()

	handleClient(clientChan)
}

func handleClient(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()

	var reply string
	err := client.Call("StatusService.GetStatus", "test", &reply)
	if err != nil {
		log.Fatal("call failed.", err)
	}
	fmt.Println(reply)
}
