package core

import (
	"errors"
	"fmt"
	"sync"
)

type GPM interface {
	Init()
	Add(IDName) error
	GetByName(string) IDName
	GetByID(int64) IDName
	Remove(IDName)
	OnRemove(IDName)
	OnAdd(IDName)
}

type BaseGPM struct {
	sync.RWMutex
	ids   map[int64]IDName
	names map[string]IDName
}

func (m *BaseGPM) Init() {
	m.ids = make(map[int64]IDName)
	m.names = make(map[string]IDName)
}

func (m *BaseGPM) Add(in IDName) error {
	m.Lock()
	defer m.Unlock()

	name := in.GetName()
	id := in.GetID()
	err := fmt.Sprintf("object already exists name: %s, id: %d", name, id)
	if _, ok := m.names[name]; ok {
		return errors.New(err)
	}

	if _, ok := m.ids[id]; ok {
		return errors.New(err)
	}
	m.names[name] = in
	m.ids[id] = in

	return nil
}

func (m *BaseGPM) GetByName(name string) IDName {
	m.RLock()
	defer m.RUnlock()

	if name != "" {
		if v, ok := m.names[name]; ok {
			return v
		}
	}
	return nil
}

func (m *BaseGPM) GetByID(id int64) IDName {
	m.RLock()
	defer m.RUnlock()

	if v, ok := m.ids[id]; ok {
		return v
	}
	return nil
}

func (m *BaseGPM) Remove(in IDName) {
	m.Lock()
	defer m.Unlock()

	name := in.GetName()
	id := in.GetID()

	if _, ok := m.names[name]; ok {
		delete(m.names, name)
	}
	if _, ok := m.ids[id]; ok {
		delete(m.ids, id)
	}
}
