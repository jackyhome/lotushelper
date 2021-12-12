package main

import (
	"flag"
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
	serverPort := flag.String("port", "", "Server port")
	flag.Parse()

	interfaces.RegisterStatusService(new(services.StatusService))
	http.HandleFunc("/jltrpc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	if *serverPort == "" {
		*serverPort = "4200"
	}
	log.Println("Server to be used: " + *serverPort)

	err := http.ListenAndServeTLS(":"+*serverPort, "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}

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
