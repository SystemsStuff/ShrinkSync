package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SystemsStuff/ShrinkSync/infra/utils"
)

func InfraHealthHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte("Method Not Allowed"))
	}
	nodes_status := utils.GetNodesStatus()
	encoder := json.NewEncoder(rw)
	if err := encoder.Encode(nodes_status); err != nil {
		http.Error(rw, "Failed to encode the nodes_status JSON response", http.StatusInternalServerError)
	}

}
