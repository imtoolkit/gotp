package main

import (
	"fmt"
	"go-otp/core"
)

func main() {
	h := core.Node{
		Port: 12345,
		Host: "localhost",
	}
	g, e := core.NewGPMD(h)
	if e != nil {
		fmt.Println("create gpmd error: ", e.Error())
		return
	}
	g.Run()
	select {}
}
