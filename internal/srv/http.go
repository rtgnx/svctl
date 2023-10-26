package srv

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rtgnx/svctl/internal/db"
	"github.com/rtgnx/svctl/internal/srv/web"
	"github.com/rtgnx/svctl/pkg/proto"
)

func GETStatus(store *db.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		msgs, err := store.ReadAll()
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, msgs)
	}
}

func POSTPush(store *db.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		msg := new(proto.Message)
		if err := c.Bind(msg); err != nil {
			c.Logger().Error(err)
			return c.NoContent(http.StatusBadRequest)
		}

		if err := store.WriteHostUpdate(*msg); err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}

func Serve(store *db.Store, cfg ServerConfig) error {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: func(origin string) (bool, error) { return true, nil },
	}))
	e.GET("/*", func(c echo.Context) error {
		assetHandler := http.FileServer(web.FS())
		return echo.WrapHandler(assetHandler)(c)
	})

	e.GET(proto.APIV1State, GETStatus(store))
	e.POST(proto.APIV1Push, POSTPush(store))

	switch cfg.UseTLS {
	case true:
		return e.StartTLS(cfg.Addr, cfg.TLSCertFile, cfg.TLSKeyFile)
	case false:
		return e.Start(cfg.Addr)
	}

	return nil
}
