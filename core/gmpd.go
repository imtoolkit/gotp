package core

import (
	"net"
	"go-otp/core/gpm"
	"fmt"
	"errors"
)

type GPMD struct {
	gpm.GPM
	gpm.Host
	hostIDMaker *gpm.IDMaker
	nodeIDMaker *gpm.IDMaker
	port        string
	ln          net.Listener
}

func NewGPMD(h *gpm.Host) (*GPMD, error) {
	ln, err := net.Listen("tcp", ":"+ h.Port)
	if err != nil {
		return nil, err
	}
	gpmd := &GPMD{
		port: h.Port,
		ln:   ln,
		hostIDMaker: gpm.NewIDMaker(0),
		nodeIDMaker: gpm.NewIDMaker(0),
	}
	gpmd.Init()
	hostName := h.Host + ":" + h.Port
	h.Name = hostName
	h.ID = 0
	gpmd.Host = *h
	gpmd.AddHost(&gpmd.Host)
	return gpmd, nil
}

func (g *GPMD) RegisterNode(n *Node) error {
	name := n.Name

	n1 := g.GetNodeByName(name)
	if n1 != nil {
		err := errors.New("can't register node in same name "+ name)
		return err
	}
	id := g.nodeIDMaker.GetInternal()
	n.ID = id
	return g.AddNode(&n.Node)
}

func (g *GPMD) RegisterHost(h *Host) error {
	name := h.Name
	h1 := g.GetHostByName(name)
	if h1 != nil {
		err := errors.New("can't register host in same name "+ name)
		return err
	}
	id := g.hostIDMaker.GetInternal()
	h.ID = id
	return g.AddHost(&h.Host)
}

func (g *GPMD)waitRegister(conn net.Conn) error {
	rpcCodec := gpm.NewRPCCodec(conn)
	rpcData := gpm.RPCData{}
	err := rpcCodec.Decode(&rpcData)
	if err != nil {
		conn.Close()
		return err
	}
	if rpcData.Type == gpm.RpcTypeHostUp {
		h := rpcData.Data.(gpm.Host)
		host := &Host{
			rpcCodec: rpcCodec,
			conn: conn,
		}
		host.Host = h
		err = g.RegisterHost(host)
		if err != nil {
			conn.Close()
			return err
		}
		fmt.Println(host.Host)
		err = host.Run(g)
	}else if rpcData.Type == gpm.RpcTypeNodeUp {
		n := rpcData.Data.(gpm.Node)
		node := &Node{
			rpcCodec: rpcCodec,
			conn:     conn,
		}
		node.Node = n
		err = g.RegisterNode(node)
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			return err
		}
		fmt.Println(node.Node)
		err = node.Run(g)
	}
	return err
}

func (g *GPMD) Run() error {
	for ;; {
		conn, err := g.ln.Accept()
		if err != nil {
			continue
		}
		go g.waitRegister(conn)
	}
	return nil
}
