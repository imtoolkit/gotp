package main

import (
	"gotp/gen_server"
)

func init() {
	gen_server.Register("test_server", func() gen_server.Server {
		return &TestServer{}
	})
}

type TestServer struct{}

func (t *TestServer) Init(args ...interface{}) gen_server.Result {
	return gen_server.ResultOk
}

func (t *TestServer) Terminate(reason gen_server.Result) gen_server.Result {
	return nil
}

func (t *TestServer) HandleCall(from gen_server.Server, request interface{}) gen_server.Result {
	return nil
}

func (t *TestServer) HandleCast(request interface{}) gen_server.Result {
	return nil
}

func (t *TestServer) HandleInfo(info interface{}) gen_server.Result {
	return nil
}

func main() {
	r, s := gen_server.Create(nil, "test_server", "test_server_1", 111)
	gen_server.Start(nil)
}
