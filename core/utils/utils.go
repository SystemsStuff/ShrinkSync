package utils

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func PingNode(node string) error {
	cmd := exec.Command("ping", "-c", "2", node)
	err := cmd.Run()
	return err
}

func GetNodesStatus() map[string]string {
	nodes := GetNodes()
	response := make(map[string]string)
	for _, node := range nodes {
		if err := PingNode(node); err != nil {
			response[node] = "DOWN"
		} else {
			response[node] = "UP"
		}
	}
	return response
}

func GetNodes() []string {
	mapNodes, reduceNodes := GetNodesCountByType()
	res := []string{}
	for i := 1; i <= mapNodes; i++ {
		res = append(res, "map-"+strconv.Itoa(int(i)))
	}
	for i := 1; i <= reduceNodes; i++ {
		res = append(res, "reduce-"+strconv.Itoa(int(i)))
	}
	res = append(res, "master")
	return res
}

func GetNodesCountByType() (int, int) {
	mapNodes, err := strconv.Atoi(os.Getenv("MAP_NODE_COUNT"))
	if err != nil {
		log.Fatal("Error parsing MAP_NODE_COUNT")
	}
	reduceNodes, err := strconv.Atoi(os.Getenv("REDUCE_NODE_COUNT"))
	if err != nil {
		log.Fatal("Error parsing REDUCE_NODE_COUNT")
	}
	return mapNodes, reduceNodes
}

func StatusCheck() {
	for {
		log.Printf("Heartbeat check...")
		response := GetNodesStatus()
		nodes := GetNodes()
		for _, node := range nodes {
			if response[node] != "UP" {
				// take appropriate action when the node is down
				log.Printf("node %s is DOWN...\n", node)
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
