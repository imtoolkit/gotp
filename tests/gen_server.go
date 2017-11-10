package main

import (
	"fmt"
	"go-otp/gen_server"
	"reflect"
)

func init() {
	gen_server.Export("test_server", TestServer{})
}

type TestServer struct {
	gen_server.ServerBase
}

func (t *TestServer) Init(args []interface{}) (result interface{}) {
	t.ServerBase.Init(args)
	gen_server.Init(&t.ServerBase, args)
	return "hello"
}

func (t *TestServer) Terminate(reason int, state interface{}) {

}

func (t *TestServer) HandleCall(request interface{}, from interface{}, state interface{}) (result interface{}) {
	fmt.Println(request, from, state)
	return
}

func (t *TestServer) HandleCast(request interface{}, state interface{}) (result interface{}) {
	return
}
func main() {
	var g = TestServer{}
	var i interface{}
	i = g
	fmt.Println(reflect.ValueOf(i).NumMethod())
	g.Init([]interface{}{})
	g.Start4("test_server_1", "main.TestServer", []interface{}{"test_server_1"}, map[string]interface{}{})
	replay := g.Call2("test_server_1", "hello")
	fmt.Println(replay)
}
