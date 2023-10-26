package mem

import (
	"encoding/json"
	"errors"
	"path"
	"sync"

	"github.com/rtgnx/svctl/internal/db/kv"
)

// KV is a fully in-memory KV store.  It is intended to be used
// with integration tests, not in production.  It is not included in
// release builds by default.
type KV struct {
	sync.RWMutex
	m map[string][]byte
}

func New() kv.KVStore {
	return &KV{sync.RWMutex{}, map[string][]byte{}}
}

// Put stores a value
func (kv *KV) Put(k string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	kv.Lock()
	kv.m[k] = b
	kv.Unlock()

	return nil
}

// Get retrives a value
func (kv *KV) Get(k string, v interface{}) error {
	kv.RLock()
	defer kv.RUnlock()
	b, ok := kv.m[k]
	if !ok {
		return errors.New("no such key")
	}
	return json.Unmarshal(b, v)
}

// Del removes a value for a given key
func (kv *KV) Delete(k string) error {
	kv.Lock()
	delete(kv.m, k)
	kv.Unlock()
	return nil
}

func (kv *KV) Keys(filter string) []string {
	kv.RLock()
	defer kv.RUnlock()

	out := []string{}
	for k := range kv.m {
		if match, _ := path.Match(filter, k); match {
			out = append(out, k)
		}
	}
	return out
}
