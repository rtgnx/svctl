package kv

type KVStore interface {
	Put(string, any) error
	Get(string, interface{}) error
	Delete(string) error
	Keys(string) []string
}
