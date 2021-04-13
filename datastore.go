package datastore

type DataStore interface {
	Get(key string) (value interface{}, err error)
	Set(key string, value interface{}) error
}
