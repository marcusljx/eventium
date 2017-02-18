package db

type DB interface {
	Put(key string, value interface{}) error
	Get(key string) (interface{}, error)
}
