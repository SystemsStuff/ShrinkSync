package master

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/SystemsStuff/ShrinkSync/core/utils"
	"github.com/jlaffaye/ftp"
)

type MasterContext struct {
	nodeStatusMap sync.Map
	mapJobsStatusMap map[string]int
	numMapTasks int
	numReducePartitions int
	numMapNodes int
	numReduceNodes int
}

func InitMaster(mux *http.ServeMux, dataGridConn *ftp.ServerConn, metadataFileName string) {

	masterContext := &MasterContext{}
	masterContext.numMapNodes, masterContext.numReduceNodes = utils.GetNodesCountByType()
	masterContext.mapJobsStatusMap = createMapJobsStatusMap(dataGridConn, metadataFileName)
	//TODO:Add logic for number of map tasks and reduce partitions
	masterContext.numMapTasks, masterContext.numReducePartitions = len(masterContext.mapJobsStatusMap),5

	fmt.Println(masterContext)
	
	masterContext.startListeners(mux)
	masterContext.startNodeStatusMonitor( 100 * time.Millisecond )
}

func (masterContext *MasterContext) startListeners(mux *http.ServeMux) {
	mux.HandleFunc("GET /infra-health", masterContext.InfraHealthHandler)
}

func (masterContext *MasterContext) startNodeStatusMonitor( interval time.Duration ) {
	// go routine to periodically check status of all worker nodes, need to handle the case of this becoming an "orphan"
	go utils.StatusCheck(&masterContext.nodeStatusMap, interval)
}

func createMapJobsStatusMap(dataGridConn *ftp.ServerConn, metadataFileName string) map[string]int {
	response, err := dataGridConn.Retr(metadataFileName)
	if err != nil {
		panic(err)
	}
	defer response.Close()
	fileBytes, err := io.ReadAll(response)
	if err != nil {
		panic(err)
	}

	jobStatusMap := make(map[string]int)
    for _, jobID := range strings.Split(string(fileBytes), "\n")  {
             jobStatusMap[jobID] = 0
    }
	return jobStatusMap
}