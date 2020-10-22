package swaggerui

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rakyll/statik/fs"
	_ "github.com/utilitywarehouse/swaggerui/statik"
)

func SwaggerUI() http.Handler {
	assetFs, err := fs.New()
	if err != nil {
		panic(fmt.Errorf("failed to create swagger-ui static assets: %w", err).Error())
	}

	return http.FileServer(assetFs)
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
