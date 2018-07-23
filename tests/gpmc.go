package main

import (
	"gotp/core/gpm"
	"os"
	"strconv"
	"strings"
)

func main() {
	names := strings.Split(os.Args[1], "@")
	if len(names) < 2 {
		return
	}
	var port = 12345
	nodeName := names[0]
	hostNames := strings.Split(names[1], ":")
	hostName := hostNames[0]
	if len(hostNames) > 1 {
		port, _ = strconv.Atoi(hostNames[1])
	}
	cfg := gpm.NodeConfig{
		Host: hostName,
		Port: port,
	}
	gn := gpm.NewGNode(cfg)
	gn.SetName(nodeName)
	gn.Connect()
	gn.Run(gn)
}
