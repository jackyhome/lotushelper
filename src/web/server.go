package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/var/www/html/helper-web")))
	log.Fatal(http.ListenAndServe(":7800", nil))
}
