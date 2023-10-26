package bitcask

import (
	"encoding/json"
	"path"
	"path/filepath"

	"git.mills.io/prologic/bitcask"
	"github.com/rtgnx/svctl/internal/db/kv"
)

type Bitcask struct {
	prefix string
	*bitcask.Bitcask
}

func New(prefix, fp string) (kv.KVStore, error) {
	bc, err := bitcask.Open(fp)
	return &Bitcask{prefix, bc}, err
}

func (bc *Bitcask) Prefix(p string) kv.KVStore {
	return &Bitcask{path.Join(bc.prefix, p), bc.Bitcask}
}

func (bc *Bitcask) Put(key string, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return bc.Bitcask.Put([]byte(key), b)
}

func (bc *Bitcask) Get(key string, v interface{}) error {

	b, err := bc.Bitcask.Get([]byte(key))
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func (bc *Bitcask) Delete(key string) error {
	return bc.Bitcask.Delete([]byte(key))
}

func (bc *Bitcask) Keys(prefix string) []string {
	keys := []string{}

	for k := range bc.Bitcask.Keys() {

		if m, _ := filepath.Match(prefix, string(k)); m {
			keys = append(keys, string(k))
		}
	}

	return keys
}
