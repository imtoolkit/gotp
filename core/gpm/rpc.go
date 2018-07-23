package gpm

import (
	"encoding/gob"
	"net"
)

const (
	RpcVersion = 1
)

type RPCData struct {
	Async   bool
	Version uint64
	ID      int64
	Data    interface{}
}

type RPCResponse struct {
	TargetID int64
	Data     interface{}
}

type RPC interface {
	OnRPC(data *RPCData) (resp interface{}, err error)
	OnAsyncRPC(data *RPCData) error
	Send(data *RPCData) error
	Receive() (*RPCData, error)
	Close() error
	OnClosed() error
	Open(conn net.Conn) error
	Run(RPC) error
}

type RPCCodec struct {
	idMaker *IDMaker
	dec     *gob.Decoder
	enc     *gob.Encoder
}

func init() {
	RegisterType(RPCData{})
	RegisterType(RPCResponse{})
}

func RegisterType(t interface{}) {
	gob.Register(t)
}

func (c *RPCCodec) Init(conn net.Conn) {
	c.idMaker = &IDMaker{}
	c.dec = gob.NewDecoder(conn)
	c.enc = gob.NewEncoder(conn)
	c.idMaker.Init(1)
}

func (c *RPCCodec) Send(data *RPCData) error {
	data.Version = RpcVersion
	data.ID = c.idMaker.Get()
	return c.enc.Encode(data)
}

func (c *RPCCodec) Receive() (*RPCData, error) {
	data := &RPCData{}
	err := c.dec.Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
