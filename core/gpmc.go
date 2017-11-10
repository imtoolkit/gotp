package core

import (
	"net"
	"go-otp/core/gpm"
)

type GPMC struct {
	gpm.GPM
	port string
	conn net.Conn
}

func NewGPMC(port string) (*GPMC, error) {
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	gpmc := &GPMC{
		port: port,
		conn: conn,
	}
	gpmc.Init()
	return gpmc, nil
}
