package db

import (
	"testing"

	"github.com/rtgnx/svctl/internal/db/mem"
	"github.com/rtgnx/svctl/pkg/proto"
	"github.com/stretchr/testify/assert"
)

func TestReadAll(t *testing.T) {

	kv := mem.New()

	store := NewStore(kv)

	msg := proto.Message{
		Hostname: "test",
	}

	assert.NoError(t, store.WriteHostUpdate(msg))

	keys := store.db.Keys("*")
	assert.GreaterOrEqual(t, len(keys), 1)
	hosts, err := store.ListHosts()

	if assert.NoError(t, err) {
		assert.Equal(t, len(hosts), 1)
	}

	msgs, err := store.ReadAll()

	if assert.NoError(t, err) {
		assert.Equal(t, len(msgs), 1)
	}

}
