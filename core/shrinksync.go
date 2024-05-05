package core

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/SystemsStuff/ShrinkSync/core/maptask"
	"github.com/SystemsStuff/ShrinkSync/core/master"
	"github.com/SystemsStuff/ShrinkSync/core/reducetask"
	"github.com/jlaffaye/ftp"
)

var MASTER = "master"
var MAP = "map"
var REDUCE = "reduce"
var DATAGRID = "datagrid"
var NETWORK = "shrink-sync-network"
var PORT = 21

type mapTask func(lineItem, lineIndex string) string
type reduceTask func(key string, valuesList []string)

type ShrinkSyncJob struct {
	mapTask               mapTask
	reduceTask            reduceTask
	inputMetadataFileName string
}

func NewShrinkSyncJob() *ShrinkSyncJob {
	return &ShrinkSyncJob{}
}

func (job *ShrinkSyncJob) SetMapTask(mapper mapTask) {
	job.mapTask = mapper
}

func (job *ShrinkSyncJob) SetReduceTask(reducer reduceTask) {
	job.reduceTask = reducer
}

func (job *ShrinkSyncJob) SetInputMetadataLocation(metadataFileName string) {
	job.inputMetadataFileName = "Input/" + metadataFileName
}

func (shrinkSyncJob *ShrinkSyncJob) Start() {
	mux := http.NewServeMux()
	dataGridConn := getDatagridConnection(DATAGRID, NETWORK, PORT)

	switch getNodeType(os.Getenv("NAME")) {
	case MASTER:
		master.InitMaster(mux, dataGridConn, shrinkSyncJob.inputMetadataFileName)
	case REDUCE:
		mux.HandleFunc("GET /endpoint", reducetask.ReduceHandler)
	case MAP:
		mux.HandleFunc("GET /endpoint", maptask.MapHandler)
	}

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", mux)
}

func getDatagridConnection(server string, network string, port int) *ftp.ServerConn {
	var err error
	var conn *ftp.ServerConn

	conn, err = ftp.Dial(fmt.Sprintf("%s.%s:%d", server, network, port), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connected to the server %s in the network %s on port %d\n", server, network, port)

	err = conn.Login("foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully logged in as user foo\n")

	return conn
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
