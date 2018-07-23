package server

import (
	"errors"
	"fmt"
	"gotp/core/gpm"
	"log"
	"net"
)

type GPMD struct {
	gpm.Node
	gpm.BaseGPM
	gpm.Conn
	conn net.Conn
	im   gpm.IDMaker
	ln   net.Listener
}

func NewGPMD(n gpm.NodeConfig) (*GPMD, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", n.Port))
	if err != nil {
		return nil, err
	}
	gpmd := &GPMD{
		ln: ln,
		im: gpm.IDMaker{},
	}
	gpmd.BaseGPM.Init()
	gpmd.NodeConfig = n
	gpmd.im.Init(0)
	gpmd.ID = gpmd.im.Get()
	gpmd.Add(gpmd)
	return gpmd, nil
}

func (g *GPMD) RegisterNode(n *gpm.Node) error {
	name := n.Name

	n1 := g.GetByName(name)
	if n1 != nil {
		err := errors.New("can't register node in same name " + name)
		return err
	}
	id := g.im.Get()
	n.ID = id
	return g.Add(g)
}

func (g *GPMD) Run() error {
	for {
		conn, err := g.ln.Accept()
		if err != nil {
			continue
		}
		n := &gpm.GNode{
			Type: gpm.NodeTypeReceiver,
		}
		n.Init(conn)
		go n.Start(g)
	}
	return nil
}

func (g *GPMD) Add(in gpm.IDName) error {
	in.SetID(g.im.Get())
	g.BaseGPM.Add(in)
	log.Println("-------on add ", in)
	return nil
}

func (g *GPMD) OnRemove(in gpm.IDName) {
}

func (g *GPMD) OnAdd(in gpm.IDName) {
}
