package kv

type KV interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}
