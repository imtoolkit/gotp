package gpm

import (
	"sync/atomic"
)

type IDName struct {
	ID   int64
	Name string
}

type IDMaker struct{
	id int64
}

func NewIDMaker(start int64) *IDMaker {
	im := &IDMaker{
		id: start - 1,
	}
	return im
}

func (im *IDMaker) GetInternal() int64 {
	return atomic.AddInt64(&im.id, 1)
}

func (im *IDMaker) GetLocal() int64 {
	return 0
}

func (im *IDMaker) GetGlobal() int64 {
	return 0
}
