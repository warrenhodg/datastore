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

// Get retrieves the value stored with the specified key
func (m *Mock) Get(key interface{}) (value interface{}, err error) {
	if m.GetError != nil {
		return nil, m.GetError
	}
	value, ok := m.Values.Load(key)
	if !ok {
		return nil, ErrNotFound
	}
	return value, nil
}

// Set sets the specified key to the specified value
func (m *Mock) Set(key interface{}, value interface{}) error {
	if m.SetError != nil {
		return m.SetError
	}
	m.Values.Store(key, value)
	return nil
}
