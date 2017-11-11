package core

import (
	"log"
	"net"
)

type Conn struct {
	RPCCodec
	conn net.Conn
}

func (c *Conn) Open(conn net.Conn) error {
	c.Init(conn)
	c.conn = conn
	return nil
}

func (c *Conn) Close() error {
	c.conn.Close()
	return nil
}

func (c *Conn) Run(r RPC) error {
	var err error
	for {
		data, err := c.Receive()
		log.Println(data)
		if err != nil {
			log.Println(err.Error())
			r.OnClosed()
			break
		} else {
			if data.Async == false {
				resp, err := r.OnRPC(data)
				if err != nil {
					log.Println(err.Error())
				} else {
					rpcResp := &RPCResponse{
						TargetID: data.ID,
						Data:     resp,
					}
					data := &RPCData{
						Async: true,
						Data:  rpcResp,
					}
					err = c.Send(data)
					if err != nil {
						log.Println(err.Error())
					}
				}
			} else {
				err = r.OnAsyncRPC(data)
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
	}
	return err
}
