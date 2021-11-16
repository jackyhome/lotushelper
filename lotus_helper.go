package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/jackyhome/lotushelper/src/interfaces"
	"github.com/jackyhome/lotushelper/src/services"
)

func main() {
	hName, _ := os.Hostname()
	fmt.Println(hName)
	// out, err := exec.Command("bash", "test.sh").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(out))

	interfaces.RegisterStatusService(new(services.StatusService))

	http.HandleFunc("/ltjsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":7999", nil)
}
