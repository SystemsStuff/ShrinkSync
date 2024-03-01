package utils

import (
	"log"
	"os"
	"strconv"
)

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
