package main

import (
	"go-otp/core"
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
	n := core.Node{
		Name: nodeName,
		Host: hostName,
		Port: port,
	}
	gn := &core.GNode{}
	gn.Connect(n)
	gn.Run(gn)
}
