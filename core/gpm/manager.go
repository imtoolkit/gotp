package gpm

import (
	"errors"
	"fmt"
	"sync"
)

type Manager struct {
	sync.RWMutex
	ids   map[int64]interface{}
	names map[string]interface{}
}

func NewManager() *Manager {
	m := &Manager{
		ids:   make(map[int64]interface{}),
		names: make(map[string]interface{}),
	}

	return m
}

func (m *Manager) Add(in IDName, o interface{}) error {
	m.Lock()
	defer m.Unlock()

	err := fmt.Sprintf("object already exists name: %s, id: %d", in.Name, in.ID)
	if _, ok := m.names[in.Name]; ok {
		return errors.New(err)
	}

	if _, ok := m.ids[in.ID]; ok {
		return errors.New(err)
	}
	m.names[in.Name] = o
	m.ids[in.ID] = o

	return nil
}

func (m *Manager) Get(in IDName) interface{} {
	m.RLock()
	defer m.RUnlock()

	if in.Name != "" {
		if v, ok := m.names[in.Name]; ok {
			return v
		}
	}
	if v, ok := m.ids[in.ID]; ok {
		return v
	}
	return nil
}

func (m *Manager) Remove(in IDName) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.names[in.Name]; ok {
		delete(m.names, in.Name)
	}
	if _, ok := m.ids[in.ID]; ok {
		delete(m.ids, in.ID)
	}
}
