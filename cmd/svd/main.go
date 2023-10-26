package main

import (
	"log"
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/rtgnx/svctl/internal/db"
	bc "github.com/rtgnx/svctl/internal/db/bitcask"
	"github.com/rtgnx/svctl/internal/srv"
)

func main() {
	app := cli.App("svd", "sv collection server")

	cfg := srv.ServerConfig{}

	app.BoolOptPtr(&cfg.UseTLS, "tls", false, "use tls")
	app.StringOptPtr(&cfg.TLSCertFile, "tls-cert", "", "tls cert file")
	app.StringOptPtr(&cfg.TLSKeyFile, "tls-key", "", "tls key file")
	app.StringOptPtr(&cfg.Addr, "addr", ":8080", "address to bind")

	var (
		dbFile = app.StringOpt("db", "./db.bitcask", "bitcask db file")
	)

	app.Action = func() {

		bc, err := bc.New("", *dbFile)

		if err != nil {
			log.Fatal(err)
		}

		if err := srv.Serve(db.NewStore(bc), cfg); err != nil {
			log.Fatal(err)
		}
	}

	app.Run(os.Args)
}
