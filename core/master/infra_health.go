package master

import (
	"encoding/json"
	"net/http"

	"github.com/SystemsStuff/ShrinkSync/core/utils"
)

func (masterContext *MasterContext) InfraHealthHandler(rw http.ResponseWriter, r *http.Request) {
	nodes_status := utils.GetNodesStatus(&masterContext.nodeStatusMap)
	rw.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(rw)
	if err := encoder.Encode(nodes_status); err != nil {
		http.Error(rw, "Failed to encode the nodes_status JSON response", http.StatusInternalServerError)
	}
}
