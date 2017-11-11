package core

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type GPMD struct {
	Node
	BaseGPM
	Conn
	conn net.Conn
	im   IDMaker
	ln   net.Listener
}

func NewGPMD(n Node) (*GPMD, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", n.Port))
	if err != nil {
		return nil, err
	}
	gpmd := &GPMD{
		ln: ln,
		im: IDMaker{},
	}
	gpmd.BaseGPM.Init()
	gpmd.Node = n
	gpmd.im.Init(0)
	gpmd.ID = gpmd.im.Get()
	gpmd.Add(gpmd)
	return gpmd, nil
}

func (g *GPMD) RegisterNode(n *Node) error {
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
		n := &GNode{
			Type: NodeTypeReceiver,
		}
		n.Init(conn)
		go n.Start(g)
	}
	return nil
}

func (g *GPMD) Add(in IDName) error {
	in.SetID(g.im.Get())
	g.BaseGPM.Add(in)
	log.Println("-------on add ", in)
	return nil
}

func (g *GPMD) OnRemove(in IDName) {
}

func (g *GPMD) OnAdd(in IDName) {
}
