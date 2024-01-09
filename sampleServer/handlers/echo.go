package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func EchoHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Fatal("Method not supported")
	}
	log.Println("Received echo request from: ", r.RemoteAddr)
	resp := fmt.Sprintf("Echo response from: %v", r.Host)
	rw.Write([]byte(resp))
}
