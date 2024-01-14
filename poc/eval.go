package main

import (
	"github.com/traefik/yaegi/interp"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/eval", evalHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func evalHandler(rw http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Error reading request body", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Raw data received data: %s\n", string(body))
	keyValuePairs := getKeyValuePairs(string(body))
	fmt.Printf("Function provided: %s\n", keyValuePairs["function"])
	fmt.Printf("Method inputs provided: %s\n", keyValuePairs["args"])

	// eval impl
	i := interp.New(interp.Options{})

	_, err = i.Eval(keyValuePairs["function"])
	if err != nil {
		log.Fatal(err)
	}

	result, err := i.Eval(keyValuePairs["args"])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Output:", result)
	response := fmt.Sprintf("Result: %v", result)
	rw.Write([]byte(response))
}

func getKeyValuePairs(s string) map[string]string {
	pairs := make(map[string]string)

	for _, pair := range strings.Split(s, "&") {
		parts := strings.Split(pair, "=")
		if len(parts) == 2 {
			key, value := parts[0], parts[1]
			pairs[key] = value
		}
	}

	return pairs
}
