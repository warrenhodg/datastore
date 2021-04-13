package datastore

import (
	"sync"
)

type Mock struct {
	Values   *sync.Map
	GetError error
	SetError error
}

func NewMock() *Mock {
	return &Mock{
		Values:   &sync.Map{},
		GetError: nil,
		SetError: nil,
	}
}

func (m *Mock) Get(key string) (value interface{}, err error) {
	if m.GetError != nil {
		return nil, m.GetError
	}
	value, ok := m.Values.Load(key)
	if !ok {
		return nil, nil
	}
	return value, nil
}

func (m *Mock) Set(key string, value interface{}) error {
	if m.SetError != nil {
		return m.SetError
	}
	m.Values.Store(key, value)
	return nil
}
