package gen_server

import (
	"fmt"
	"go-otp/core"
	"reflect"
	"time"
)

const (
	ReasonNormal = iota
	ReasonShutdown
	OptNormal
	OptTerminate
	ResultOk
	ResultError
)

type Result int

var g_type_refs = make(map[string]interface{})
var g_server_refs = make(map[string]*ServerBase)

var g_server_root = ServerBase{}

type Server interface {
	Init(args []interface{}) (result interface{})
	Terminate(reason int, state interface{})
	HandleCall(request interface{}, from interface{}, state interface{}) (result interface{})
	HandleCast(request interface{}, state interface{}) (result interface{})
}

type ServerBase struct {
	core.Server
}

func init() {
}

func (g *ServerBase) loop() {
}

func registerServer(name string, g *ServerBase) {
	g_server_refs[name] = g
}

func getServer(serverRef string) *ServerBase {
	return g_server_refs[serverRef]
}

func registerType(name string, g interface{}) {
	g_type_refs[name] = g
}

func getType(module string) interface{} {
	return g_type_refs[module]
}

func Export(name string, g interface{}) {
	registerType(name, g)
}

func Init(g *ServerBase, args []interface{}) (result interface{}) {
	g.c = make(chan interface{})

	go g.loop()

	return
}

func (g *ServerBase) init(name string, args []interface{}) {
}

func doAbCast(from *ServerBase, nodes []string, serverName string, request interface{}) {

}

func doCall(from *ServerBase, serverRef string, request interface{}, timeout time.Duration) (replay interface{}) {
	return
}

func doStart(from *ServerBase, serverName, module string, args []interface{}, options map[string]interface{}) (result interface{}) {
	server := getType(module)
	fmt.Println(reflect.ValueOf(server).NumMethod())

	ret := reflect.ValueOf(server).MethodByName("Init").Call([]reflect.Value{reflect.ValueOf(args)})
	return ret[0]
}

//--------------------
func (g *ServerBase) AbCast2(name string, request interface{}) {
	g.AbCast3([]string{}, name, request)
}

func (g *ServerBase) AbCast3(nodes []string, name string, request interface{}) {
	doAbCast(g, nodes, name, request)
}

func (g *ServerBase) Call2(serverRef string, request interface{}) (replay interface{}) {
	return g.Call3(serverRef, request, 0)
}

func (g *ServerBase) Call3(serverRef string, request interface{}, timeout time.Duration) (replay interface{}) {
	return doCall(g, serverRef, request, timeout)
}

func (g *ServerBase) EnterLoop4(model string, options map[string]interface{},
	state interface{}, timeout time.Duration) {

}

func (g *ServerBase) EnterLoop5(model string, options map[string]interface{},
	state interface{}, serverName string, timeout time.Duration) {

}

func (g *ServerBase) MultiCall2(name string, request interface{}) (result interface{}) {
	return
}

func (g *ServerBase) MultiCall3(nodes []string, name string, request interface{}) (result interface{}) {
	return
}

func (g *ServerBase) MultiCall4(nodes []string, name string, request interface{}, timeout time.Duration) (result interface{}) {
	return
}

func (g *ServerBase) Reply2(client interface{}, replay interface{}) (result interface{}) {
	return
}

func (g *ServerBase) Start3(module string, args []interface{}, options map[string]interface{}) (result interface{}) {
	serverName := module
	return g.Start4(serverName, module, args, options)
}

func (g *ServerBase) Start4(serverName string, module string, args []interface{},
	options map[string]interface{}) (result interface{}) {
	return doStart(g, serverName, module, args, options)
}

func (g *ServerBase) StartLink3(module string, args []interface{}, options map[string]interface{}) (result interface{}) {
	return
}

func (g *ServerBase) StartLink4(serverName string, module string, args map[string]interface{},
	options map[string]interface{}) (result interface{}) {
	return
}

func (g *ServerBase) Stop(server interface{}) {

}

func (g *ServerBase) Stop3(server interface{}, reason string, timeout time.Duration) {

}
