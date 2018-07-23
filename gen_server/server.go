package gen_server

import (
	"sync"
	"time"
)

type Server interface {
	Init(args ...interface{}) Result
	Terminate(reason Result) Result
	HandleCall(from Server, request interface{}) Result
	HandleCast(request interface{}) Result
	HandleInfo(info interface{}) Result
}

type EventType string

const (
	EventCall      EventType = "call"
	EventCast                = "cast"
	EventTeminate            = "terminate"
	EventTeminated           = "terminated"
)

type Event struct {
	Type EventType
	From Server
	Data interface{}
}

func NewEvent(sender Server, tp EventType, data interface{}) *Event {
	return &Event{
		Type: tp,
		From: sender,
		Data: data,
	}
}

type serverWrapper struct {
	Server
	children []Server
	name     string
	module   string
	event    chan *Event
}

var manager *serverManager = nil

type serverManager struct {
	sync.RWMutex
	Server
	modules map[string]Creator
	servers map[string]Server
}

func init() {
	manager = &serverManager{
		modules: make(map[string]Creator),
		servers: make(map[string]Server),
	}
}

func (m *serverManager) registerModule(module string, c Creator) Result {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.modules[module]; ok {
		return ResultError
	}
	m.modules[module] = c
	return ResultOk
}

func (m *serverManager) createServer(module, name string, args ...interface{}) (Result, Server) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.servers[name]; ok {
		return ResultError, nil
	}
	if create, ok := m.modules[module]; ok {
		s := create()
		r := s.Init(args...)
		if r == ResultOk {
			rs := &serverWrapper{
				Server: s,
				module: module,
				name:   name,
				event:  make(chan *Event, 1000),
			}
			m.servers[name] = rs
		}
		return r, s
	}
	return ResultError, nil
}

type Creator func() Server

func Register(module string, c Creator) Result {
	return manager.registerModule(module, c)
}

func doCreate(from Server, module, name string, args ...interface{}) (Result, Server) {
	return manager.createServer(module, name, args...)
}

func Start(module, name string, options map[string]interface{}, args ...interface{}) Result {
	return StartLink(nil, module, name, options, args...)
}

func StartLink(link Server, module, name string, options map[string]interface{}, args ...interface{}) Result {
	return nil
}

func Cast(from Server, nodes []string, serverName string, request interface{}) {

}

func Call(from Server, serverName string, request interface{}, timeout ...time.Duration) Result {
	return nil
}

func MultiCall(from Server, nodes []string, module string, request interface{}, timeout ...time.Duration) Result {
	return nil
}

func Stop(from Server, s Server, reason string, timeout ...time.Duration) {

}

func Reply(from Server, to Server, replay interface{}) Result {
	return nil
}
