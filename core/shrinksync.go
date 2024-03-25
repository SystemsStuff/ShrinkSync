package core

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/SystemsStuff/ShrinkSync/core/maptask"
	"github.com/SystemsStuff/ShrinkSync/core/master"
	"github.com/SystemsStuff/ShrinkSync/core/reducetask"
)

var MASTER = "master"
var MAP = "map"
var REDUCE = "reduce"

func NewShrinkSyncJob() {
	mux := http.NewServeMux()
	
	switch getNodeType(os.Getenv("NAME")) {
	case MASTER:
		mux.HandleFunc("/infraHealth", master.MasterHandler)
	case REDUCE:
		mux.HandleFunc("/endpoint", reducetask.ReduceHandler)
	case MAP:
		mux.HandleFunc("/endpoint", maptask.MapHandler)
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
