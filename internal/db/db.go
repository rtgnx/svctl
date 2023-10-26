package db

import (
	"time"

	"github.com/rtgnx/svctl/internal/db/kv"
	"github.com/rtgnx/svctl/pkg/proto"
)

const (
	DiscardAfter time.Duration = time.Minute * 2
	Prefix                     = "runit"
)

type Store struct {
	db kv.KVStore
}

func NewStore(kv kv.KVStore) *Store {
	return &Store{db: kv}
}

func (store *Store) WriteHostUpdate(msg proto.Message) error {
	return store.db.Put(msg.Hostname, msg)
}

func (store *Store) ReadAll() ([]proto.Message, error) {
	messages := make([]proto.Message, 0)

	for _, key := range store.db.Keys("*") {
		msg := new(proto.Message)

		if err := store.db.Get(key, msg); err != nil {
			return messages, err
		}

		// if time.Since(msg.Timestamp) > DiscardAfter {
		// 	// Delete host data that hasn't been updated longer than Discard time
		// 	store.db.Delete(key)
		// 	continue
		// }

		messages = append(messages, *msg)
	}

	return messages, nil
}

func (store *Store) ListHosts() ([]string, error) {
	hosts := []string{}
	for _, key := range store.db.Keys("*") {
		hosts = append(hosts, key)
	}

	return hosts, nil
}
