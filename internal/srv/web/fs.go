package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed vue/dist
var embededFiles embed.FS

func FS() http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "vue/dist")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
