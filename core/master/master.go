package master

import (
	"fmt"
	"net/http"

	"github.com/SystemsStuff/ShrinkSync/core/utils"
)

func InitMaster(mux *http.ServeMux) {
	mapNodeCount, reduceNodeCount := utils.GetNodesCountByType()

	fmt.Println(mapNodeCount, reduceNodeCount)
	mux.HandleFunc("GET /infra-health", InfraHealthHandler)

	startNodeStatusMonitor()
}

func startNodeStatusMonitor() {
	// go routine to periodically check status of all worker nodes, need to handle the case of this becoming an "orphan"
	go utils.StatusCheck()
}
