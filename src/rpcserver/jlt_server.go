package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/jackyhome/lotushelper/src/interfaces"
	"github.com/jackyhome/lotushelper/src/services"
)

func main() {

	interfaces.RegisterStatusService(new(services.StatusService))
	http.HandleFunc("/jltrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":4200", nil)

}

func clientReq() {
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
			log.Println("Connection from: ", conn.LocalAddr())

			clientChan <- rpc.NewClient(conn)
			handleClient(clientChan)
		}
	}()
}

func handleClient(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()

	var reply []byte
	err := client.Call("StatusService.GetStatus", "test", &reply)
	if err != nil {
		log.Fatal("call failed.", err)
	}
	fmt.Println(string(reply))
}
