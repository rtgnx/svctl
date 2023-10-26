package srv

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rtgnx/svctl/internal/db"
	"github.com/rtgnx/svctl/internal/db/mem"
	"github.com/rtgnx/svctl/pkg/proto"
	"github.com/stretchr/testify/assert"
)

func TestPOSTPush(t *testing.T) {
	e := echo.New()
	msg := proto.Message{Hostname: "abc"}
	b, _ := json.Marshal(msg)
	req := httptest.NewRequest(http.MethodPost, proto.APIV1Push, strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	kv := mem.New()
	store := db.NewStore(kv)
	if assert.NoError(t, POSTPush(store)(c)) {
		storedMsg := new(proto.Message)
		assert.NoError(t, kv.Get(path.Join(db.Prefix, msg.Hostname), storedMsg))
		assert.Equal(t, http.StatusCreated, rec.Code)

		assert.Equal(t, storedMsg.Hostname, msg.Hostname)

	}
}

func TestGETState(t *testing.T) {

}
