package core

import (
	"go-otp/core/gpm"
	"net"
	"fmt"
)

type NodeArgs struct {
	Name string
	Host string
	Port string
}

type Node struct {
	gpm.Node
	rpcCodec *gpm.RPCCodec
	conn net.Conn
}

func (n *Node) Start(args *NodeArgs) error {
	hostName := args.Host + ":" + args.Port
	conn, err := net.Dial("tcp", hostName)
	if err != nil {
		return err
	}
	n.Host = gpm.Host{
		Host: args.Host,
		Port: args.Port,
	}
	n.Host.Name = hostName
	n.Host.ID = 0
	n.Name = args.Name
	n.conn = conn
	n.rpcCodec = gpm.NewRPCCodec(conn)
	rpcData := gpm.RPCData{
		Type: gpm.RpcTypeNodeUp,
		Version: 1,
		Data: n.Node,
	}
	err = n.rpcCodec.Encode(&rpcData)
	fmt.Println(err)
	n.Run(nil)
	return nil
}

func (n *Node) Run(g *GPMD) error {
	var err error = nil
	for ;; {
		rpcData := gpm.RPCData{}
		err = n.rpcCodec.Decode(&rpcData)
		fmt.Println(rpcData)
		fmt.Println(err)
		if err != nil {
			if g != nil {
				g.RemoveNode(&n.Node)
			}
			n.conn.Close()
			break
		}
	}
	return err
}
