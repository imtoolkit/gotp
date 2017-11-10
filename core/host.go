package core

import (
	"go-otp/core/gpm"
	"net"
	"fmt"
)

type Host struct {
	gpm.Host
	conn net.Conn
	rpcCodec *gpm.RPCCodec
}

func (h *Host)Run(g *GPMD) error {
	var err error = nil
	for ;; {
		rpcData := gpm.RPCData{}
		err = h.rpcCodec.Decode(&rpcData)
		fmt.Println(rpcData)
		fmt.Println(err)
	}
	return err
}
