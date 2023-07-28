package swaggerui

import (
	"fmt"
	"github.com/utilitywarehouse/swaggerui/initialiser"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/utilitywarehouse/swaggerui/static"
)

func SwaggerUI() http.Handler {
	mux := http.NewServeMux()
	handler, err := initialiser.ConfigLoader("swagger.json")
	if err != nil {
		panic("failed to create FS for swaggerui")
	}
	mux.Handle("/swagger-initializer.js", handler)
	mux.Handle("/", http.FileServer(http.FS(static.Assets)))
	return mux
}

func SwaggerUIWithSwaggerLocation(location string) (http.Handler, error) {
	mux := http.NewServeMux()
	handler, err := initialiser.ConfigLoader(location)
	if err != nil {
		return nil, fmt.Errorf("failed to create FS: %w", err)
	}
	mux.Handle("/swagger-initializer.js", handler)
	mux.Handle("/", http.FileServer(http.FS(static.Assets)))
	return mux, nil
}

func SwaggerFile(swaggerLocations ...string) http.Handler {
	location := withDefaultFileLocations(swaggerLocations)
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if location == "" {
			http.Error(writer, "no swagger file specified", http.StatusNotFound)
		}
		http.ServeFile(writer, request, location)
	})
}

func withDefaultFileLocations(locations []string) string {

	if len(locations) == 0 {
		locations = getDefaultLocations()
	}

	for _, location := range locations {
		if _, err := os.Stat(location); err == nil {
			log.Println(fmt.Sprintf("loaded swagger file at %s", location))
			return location
		}
		log.Println(fmt.Sprintf("failed to load swagger file at %s", location))
	}

	if len(locations) > 0 {
		panic("no swagger file located, either specify the location of a swagger file place one " +
			"in /tmp/swagger.json or in the same directory as the application binary")
	}

	return ""
}

func getDefaultLocations() []string {
	var locations []string
	if e, err := os.Executable(); err == nil {
		locations = append(locations, fmt.Sprintf("%s/swagger.json", filepath.Dir(e)))
	}
	locations = append(locations, "/tmp/swagger.json")
	return locations
}
