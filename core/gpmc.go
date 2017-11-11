package core

import (
	"net"
)

type GPMC struct {
	BaseGPM
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
