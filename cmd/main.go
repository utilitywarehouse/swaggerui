package main

import (
	"embed"
	"net/http"
	"time"

	"github.com/utilitywarehouse/swaggerui"
)

//go:embed swagger
var swaggerFS embed.FS

func main() {
	m := http.NewServeMux()

	m.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", swaggerui.SwaggerUI()))
	m.Handle("/swagger.json", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFileFS(writer, request, swaggerFS, "swagger/swagger.json")
	}))

	svr := &http.Server{
		Addr:              "localhost:8080",
		Handler:           m,
		ReadHeaderTimeout: 5 * time.Second,
	}

	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
}
