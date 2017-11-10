package core

import (
	"go-otp/core/gpm"
	"time"
)

type PID struct {
	Host    gpm.IDName
	Node    gpm.IDName
	Process gpm.IDName
}

type Event struct {
	PID
	To      *PID
	Name    string
	Data    interface{}
	Chan    []chan interface{}
	Timeout time.Duration
}

type Server struct {
	PID
	name   string
	c      chan interface{}
	sub    []*Server
	parent *Server
}

func StartServer(server *Server) *PID {
	return nil
}
