package utils

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

func PingNode(node string) error {
	cmd := exec.Command("ping", "-c", "1", node)
	err := cmd.Run()
	return err
}

func GetNodesStatus(statusMap *sync.Map) map[string]string {
	// Convert sync.map to map
	response := make(map[string]string)
	statusMap.Range(func(key, value any) bool {
		response[key.(string)] = value.(string)
		return true
	})

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

func updateNodeStatus(statusMap *sync.Map) {
	nodes := GetNodes()

	var wg sync.WaitGroup
	// Check status of nodes concurrently
	for _, node := range nodes {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := PingNode(node); err != nil {
				statusMap.Store(node, "DOWN")
			} else {
				statusMap.Store(node, "UP")
			}
		}()
	}
	wg.Wait()
}

func StatusCheck(statusMap *sync.Map, interval time.Duration) {
	for {
		log.Printf("Heartbeat check...")
		updateNodeStatus(statusMap)
		statusMap.Range(func(node, status any) bool {
			if status.(string) != "UP" {
				// take appropriate action when the node is down
				log.Printf("node %s is DOWN...\n", node.(string))
			}
			return true
		})
		time.Sleep(interval)
	}
}
