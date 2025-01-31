# Swagger UI

 `go get github.com/utilitywarehouse/swaggerui`

A wrapper around https://github.com/swagger-api/swagger-ui which is served a static file from a Go API. The swagger UI will be available on `<YOUR_DOMAIN>/swagger-ui/` , and will look for the `swagger.json` at `<YOUR_DOMAIN>/swagger.json` . This allows you to pass in an auth token, which will be necessary to hit many endpoints. You can get your okta token by running

```
uw iam login --copy --print --format json --quiet
```

with the UW CLI (https://github.com/utilitywarehouse/uw-cli).

## Usage

This is intended to be used alongside a Go API as following.

```
	import "github.com/utilitywarehouse/swaggerui"

	
	m := http.NewServeMux()

	swaggerFilePath := "swagger/swagger.json"

	m.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", swaggerui.SwaggerUI()))
	m.Handle("/swagger.json", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFileFS(writer, request, swaggerFS, swaggerFilePath)
	}))

	svr := &http.Server{
		Addr:              "localhost:8080",
		Handler:           m,
		ReadHeaderTimeout: 5 * time.Second,
	}

	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}
```

## Development

`cmd/swagger` contains a sample swagger file for testing purposes.

`cmd/main.go` will spin up a go server on localhost:8080 which can be used for testing out changed to the swagger ui.

Make sure to run `make build-ui` before pushing up if you're making changes to the UI; this is intended to work by serving the static files made from this UI, rather than the UI itself being run in a docker container or something like that.
