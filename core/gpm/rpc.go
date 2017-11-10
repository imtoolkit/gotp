package gpm

import (
	"encoding/gob"
	"net"
)

const (
	RpcTypeHostUp = iota
	RpcTypeNodeUp
	RpcTypeNodeDown
	RpcTypeRequest
	RpcTypeResponse
)

type RPCData struct {
	Type    int
	Version uint64
	Data    interface{}
}

type RPCCodec struct {
	dec *gob.Decoder
	enc *gob.Encoder
}

func init() {
	RegisterType(Host{})
	RegisterType(Node{})
	RegisterType(RPCData{})
}

func RegisterType(t interface{}) {
	gob.Register(t)
}

func NewRPCCodec(conn net.Conn) *RPCCodec {
	c := &RPCCodec{
		dec: gob.NewDecoder(conn),
		enc: gob.NewEncoder(conn),
	}
	return c
}

func (c *RPCCodec) Encode(data *RPCData) error {
	return c.enc.Encode(data)
}

func (c *RPCCodec) Decode(data *RPCData) error {
	return c.dec.Decode(data)
}
