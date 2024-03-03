package core

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/SystemsStuff/ShrinkSync/core/node_handlers"
)

var MASTER = "master"
var MAP = "map"
var REDUCE = "reduce"

func NewShrinkSyncJob() {
	mux := http.NewServeMux()
	
	switch getNodeType(os.Getenv("NAME")) {
	case MASTER:
		mux.HandleFunc("/infraHealth", node_handlers.MasterHandler)
	case REDUCE:
		mux.HandleFunc("/endpoint", node_handlers.ReduceHandler)
	case MAP:
		mux.HandleFunc("/endpoint", node_handlers.MapHandler)
	}

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", mux)
}

func getNodeType(name string) string {
	if strings.HasPrefix(name, MAP) {
		return MAP
	} else if strings.HasPrefix(name, REDUCE) {
		return REDUCE
	} else {
		return MASTER
	}
}
