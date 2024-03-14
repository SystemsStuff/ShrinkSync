package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/SystemsStuff/ShrinkSync/infra/handlers"
	"github.com/SystemsStuff/ShrinkSync/infra/utils"
)

func main() {

	// go routine to periodically check status of all worker nodes, need to handle the case of this becoming an "orphan"
	go utils.StatusCheck()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path[1:]))
	})
	mux.HandleFunc("/infra-health", handlers.InfraHealthHandler)

	log.Fatal(http.ListenAndServe(":9090", mux))
}
