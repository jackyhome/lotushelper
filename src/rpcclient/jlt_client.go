package main

import (
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/jackyhome/lotushelper/src/services"
)

func main() {
	rpc.Register(new(services.StatusService))

	for {
		conn, _ := net.Dial("tcp", "fc.pikumao.com:2021")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		log.Println("connect succesfully.")
		rpc.ServeConn(conn)
		conn.Close()
	}
}
