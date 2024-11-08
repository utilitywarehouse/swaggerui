package swaggerui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed app/dist/*
var assetsFS embed.FS

func SwaggerUI() http.Handler {
	assets, err := fs.Sub(assetsFS, "app/dist")
	if err != nil {
		panic(err)
	}

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.FileServer(http.FS(assets)).ServeHTTP(writer, request)
	})
}
