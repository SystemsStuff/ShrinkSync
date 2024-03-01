package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/SystemsStuff/ShrinkSync/infra/handlers"
	"github.com/SystemsStuff/ShrinkSync/infra/utils"
)

func main() {

	go func() {
		for {
			log.Printf("Heartbeat check...")
			response := utils.GetNodesStatus()
			nodes := utils.GetNodes()
			for _, node := range nodes {
				if response.StatusJson[node] != "UP" {
					// take appropriate action when the node is down
					log.Printf("node %s is DOWN...\n", node)
				}
			}
			time.Sleep(1 * time.Minute)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path[1:]))
	})
	mux.HandleFunc("/infra-health", handlers.InfraHealthHandler)

	log.Fatal(http.ListenAndServe(":9090", mux))
}
