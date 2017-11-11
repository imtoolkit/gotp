package core

import (
	"fmt"
	"net"
)

const (
	NodeTypeNormal = iota
	NodeTypeReceiver
	NodeTypeHost
)

type Node struct {
	Type int
	Name string
	Host string
	IP   string
	Port int
	ID   int64
}

const (
	RPCNodeUp = iota
	RPCNodeDown
)

type RPCNodeStatus struct {
	Status int
	Node   Node
}

func init() {
	RegisterType(Node{})
	RegisterType(RPCNodeStatus{})
}

func (n *Node) GetID() int64 {
	return n.ID
}

func (n *Node) GetName() string {
	return n.Name
}

func (n *Node) SetID(id int64) {
	n.ID = id
}

func (n *Node) SetName(name string) {
	n.Name = name
}

type NodeArgs struct {
	Node Node
	Conn net.Conn
}

type GNode struct {
	Node
	Conn
	gpm  GPM
	Type int
}

func (n *GNode) Connect(node Node) error {
	n.Node = node
	addr := fmt.Sprintf("%s:%d", n.IP, n.Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	n.Init(conn)
	data := &RPCData{
		Async: false,
		Data: &RPCNodeStatus{
			Status: RPCNodeUp,
			Node:   node,
		},
	}
	n.Send(data)
	return nil
}

func (n *GNode) Start(gpm GPM) error {
	n.gpm = gpm
	err := n.Run(n)
	return err
}

func (n *GNode) OnRPC(data *RPCData) (resp interface{}, err error) {
	switch v := data.Data.(type) {
	case RPCNodeStatus:
		n.Node = v.Node
		n.gpm.Add(n)
	}
	return nil, nil
}

func (n *GNode) OnAsyncRPC(data *RPCData) error {
	return nil
}

func (n *GNode) OnClosed() error {
	return nil
}
