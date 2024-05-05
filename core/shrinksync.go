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

type mapTask func(lineItem, lineIndex string)
type reduceTask func(key string, valuesList []string)

type ShrinkSyncJob struct {
	mapTask mapTask
	reduceTask reduceTask
}

func NewShrinkSyncJob() (*ShrinkSyncJob) {
	return &ShrinkSyncJob{};
}

func (job *ShrinkSyncJob) SetMapTask(mapper mapTask) {
	job.mapTask = mapper
}

func (job *ShrinkSyncJob) SetReduceTask(reducer reduceTask) {
	job.reduceTask = reducer
}

func (*ShrinkSyncJob) Start() {
	mux := http.NewServeMux()
	
	switch getNodeType(os.Getenv("NAME")) {
	case MASTER:
		master.InitMaster(mux)
	case REDUCE:
		mux.HandleFunc("GET /endpoint", reducetask.ReduceHandler)
	case MAP:
		mux.HandleFunc("GET /endpoint", maptask.MapHandler)
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
