package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func DiscoverNodes(network string) []string {
	cmd := fmt.Sprintf("CONTAINER_NAMES=$(curl --unix-socket /var/run/docker.sock http://host.docker.internal/containers/json | jq -r --arg NETWORK '%s' '.[] | select(.NetworkSettings.Networks[$NETWORK]) | .Names[]' | sed 's/\\///g') && echo $CONTAINER_NAMES", network)
	output, err := execute(cmd)
	if err != nil {
		log.Fatal(err)
	}
	nodes := strings.Fields(output)
	return nodes
}

func execute(cmd string) (string, error) {
	command := exec.Command("sh", "-c", cmd)
	var output bytes.Buffer
	command.Stdout = &output
	err := command.Run()
	return output.String(), err
}
