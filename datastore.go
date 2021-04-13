package datastore

import "errors"

// Errors
const (
	ErrNotFound = errors.New("Not found")
)

// GetSetter is used to get or set values in a datastore
type GetSetter interface {
	Get(key interface{}) (value interface{}, err error)
	Set(key interface{}, value interface{}) error
}
