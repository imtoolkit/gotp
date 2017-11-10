package main

import (
	"go-otp/core"
	"fmt"
	"go-otp/core/gpm"
)

func main(){
	h := &gpm.Host{
		Port: "12345",
		Host: "localhost",
	}
	g, e := core.NewGPMD(h)
	if e != nil {
		fmt.Println("create gpmd error: ", e.Error())
		return
	}
	g.Run()
	select{}
}
