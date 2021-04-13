package datastore

import "errors"

// Errors
const (
	ErrNotFound = errors.New("Not found")
)

type DataStore interface {
	Get(key interface{}) (value interface{}, err error)
	Set(key interface{}, value interface{}) error
}
