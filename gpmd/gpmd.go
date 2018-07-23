package main

import (
	"fmt"
	"gotp/core/gpm"
	"gotp/core/server"
)

func main() {
	h := gpm.NodeConfig{
		Port: 12345,
		Host: "localhost",
	}
	g, e := server.NewGPMD(h)
	if e != nil {
		fmt.Println("create gpmd error: ", e.Error())
		return
	}
	g.Run()
	select {}
}
