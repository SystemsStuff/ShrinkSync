package master

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/SystemsStuff/ShrinkSync/core/utils"
)

type MasterContext struct {
	statusMap sync.Map
	numMapTasks int
	numReducePartitions int
	numMapNodes int
	numReduceNodes int
}

func InitMaster(mux *http.ServeMux) {

	masterContext := &MasterContext{}
	masterContext.numMapNodes, masterContext.numReduceNodes = utils.GetNodesCountByType()
	// Add logic for number of map tasks and reduce partitions
	masterContext.numMapTasks, masterContext.numReducePartitions = 5,5
	fmt.Println(masterContext)
	mux.HandleFunc("GET /infra-health", masterContext.InfraHealthHandler)

	masterContext.startNodeStatusMonitor( 100 * time.Millisecond )
}

func (masterContext *MasterContext) startNodeStatusMonitor( interval time.Duration ) {
	// go routine to periodically check status of all worker nodes, need to handle the case of this becoming an "orphan"
	go utils.StatusCheck(&masterContext.statusMap, interval)
}
