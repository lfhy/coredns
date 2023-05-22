package resource

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed *
var Resource embed.FS

func GetStaticFS() http.FileSystem {
	sfs, err := fs.Sub(Resource, "static")
	if err != nil {
		panic(err)
	}
	return http.FS(sfs)
}
