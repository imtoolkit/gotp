package gpm

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
	BaseIDName
	NodeConfig
}

type NodeConfig struct {
	Type int
	Host string
	IP   string
	Port int
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

func NewGNode(cfg NodeConfig) *GNode {
	node := &GNode{}
	node.NodeConfig = cfg
	return node
}

func (n *GNode) Connect() error {
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
			Node:   n.Node,
		},
	}
	n.Send(data)
	return nil
}

func (n *GNode) Start(g GPM) error {
	n.gpm = g
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
