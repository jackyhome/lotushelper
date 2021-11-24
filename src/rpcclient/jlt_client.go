package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/jackyhome/lotushelper/src/services"
)

func main() {
	rpc.Register(new(services.StatusService))
	serverHost := flag.String("server", "", "Server host name")
	serverPort := flag.String("port", "", "Server port")
	flag.Parse()

	log.Println("Server to be used: " + *serverHost + ":" + *serverPort)
	for {
		conn, _ := net.Dial("tcp", *serverHost+":"+*serverPort)
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		log.Println("connect succesfully.")
		rpc.ServeConn(conn)
		conn.Close()
	}
}
