package main

import (
	"go-otp/core"
	"fmt"
)

func main(){
	n := &core.Node{}
	args := &core.NodeArgs{
		Name: "test_gpmc1",
		Port: "12345",
		Host: "localhost",
	}
	err := n.Start(args)
	fmt.Println(err)
}