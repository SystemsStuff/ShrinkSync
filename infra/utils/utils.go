package utils

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

type NodesStatus struct {
	StatusJson map[string]string `json:"nodes_status"`
}

func PingNode(node string) error {
	cmd := exec.Command("ping", "-w", "2", node)
	err := cmd.Run()
	return err
}

func GetNodesStatus() *NodesStatus {
	nodes := GetNodes()
	response := &NodesStatus{
		StatusJson: make(map[string]string),
	}
	for _, node := range nodes {
		if err := PingNode(node); err != nil {
			response.StatusJson[node] = "DOWN"
		} else {
			response.StatusJson[node] = "UP"
		}
	}
	return response
}

func GetNodes() []string {
	mapNodes, _ := strconv.ParseInt(os.Getenv("MAP_NODE_COUNT"), 10, 64)
	reduceNodes, err := strconv.ParseInt(os.Getenv("REDUCE_NODE_COUNT"), 10, 64)
	if err != nil {
		log.Fatal("Error parsing the env variables")
	}
	res := []string{}
	var i int64
	for i = 1; i <= mapNodes; i++ {
		res = append(res, "map-"+strconv.Itoa(int(i)))
	}
	for i = 1; i <= reduceNodes; i++ {
		res = append(res, "reduce-"+strconv.Itoa(int(i)))
	}
	res = append(res, "master")
	return res
}
