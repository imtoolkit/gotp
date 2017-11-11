package core

import (
	"sync/atomic"
)

type IDName interface {
	GetID() int64
	GetName() string
	SetID(int64)
	SetName(string)
}

type BaseIDName struct {
	ID   int64
	Name string
}

func (in *BaseIDName) GetID() int64 {
	return in.ID
}

func (in *BaseIDName) GetName() string {
	return in.Name
}

func (in *BaseIDName) SetID(id int64) {
	in.ID = id
}

func (in *BaseIDName) SetName(name string) {
	in.Name = name
}

type IDMaker struct {
	id int64
}

func (im *IDMaker) Init(start int64) {
	im.id = start - 1
}

func (im *IDMaker) Get() int64 {
	return atomic.AddInt64(&im.id, 1)
}
