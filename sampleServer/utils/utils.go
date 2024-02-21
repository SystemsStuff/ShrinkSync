package utils

import (
	"log"
	"os"
	"strconv"
)

// func DiscoverNodes(network string) []string {
// 	cmd := fmt.Sprintf("CONTAINER_NAMES=$(curl --unix-socket /var/run/docker.sock http://host.docker.internal/containers/json | jq -r --arg NETWORK '%s' '.[] | select(.NetworkSettings.Networks[$NETWORK]) | .Names[]' | sed 's/\\///g') && echo $CONTAINER_NAMES", network)
// 	output, err := execute(cmd)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	nodes := strings.Fields(output)
// 	return nodes
// }

// func execute(cmd string) (string, error) {
// 	command := exec.Command("sh", "-c", cmd)
// 	var output bytes.Buffer
// 	command.Stdout = &output
// 	err := command.Run()
// 	return output.String(), err
// }

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
