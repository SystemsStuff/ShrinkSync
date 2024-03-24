package master

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/SystemsStuff/ShrinkSync/core/utils"
)

const (
	networkName string = "shrink-sync-network"
)

type message struct {
	Msg string `json:"message"`
}

func HealthHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body message
	err := decoder.Decode(&body)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	fmt.Printf("Received post request on health handler with message: %s\n", body)
	fmt.Printf("Broadcasting /echo request to all the containers on the network...\n")
	nodes := utils.GetNodes()
	for _, node := range nodes {
		wg.Add(1)
		go func(nd string) {
			sendEchoRequest(fmt.Sprintf("http://%s:8080/echo", nd))
			wg.Done()
		}(node)
	}
	wg.Wait()
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Server is up. Sent an echo to all the nodes on the network..."))
}

func sendEchoRequest(url string) {
	fmt.Printf("Sending echo to %s...\n", url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received response from %s: %s\n", url, string(body))
}
