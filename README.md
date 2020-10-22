# Swagger UI

`go get github.com/utilitywarehouse/swaggerui`

swaggerui returns a handler to serve a Swagger API typically used in conjunction with the grpc-gateway

### Updating the UI to the latest version
simples `make update`. This will find and download the latest tagged release of swagger UI, convert the static assets 
to a binary file and clean up after itself.

### swagger.json location
you can specify the potential locations of swagger files by passing them to `SwaggerFile`, it will pick the first available
location!. Passing no locations will default to the locations of `YOUR_BINARY_LOCATION/swagger.json` then `/tmp/swagger.json`.

***Please note***
If you specify locations to `SwaggerFile` and it fails to locate a valid file the method will panic!

### Example usage

##### Overriding the default swagger.json location

```go
func initialiseSwaggerAPI(ctx context.Context, restPort, grpcPort, maxMessageSize *int, swaggerFileLocation *string) *http.Server {
	m := http.NewServeMux()

	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{},
	}))

	m.Handle("/", gwmux)
	m.Handle("/swagger-ui/", http.stripPrefix("/swagger-ui/", swaggerui.SwaggerUI()))
	m.Handle("/swagger.json", swaggerui.SwaggerFile(*swaggerFileLocation))

	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(*maxMessageSize*1024*1024),
			grpc.MaxCallSendMsgSize(*maxMessageSize*1024*1024),
		),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer())),
		)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer())),
		)),
	}

	if rErr := my_gateway.RegisterServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("localhost:%d", *grpcPort), dialOpts); rErr != nil {
		log.WithError(rErr).Panic("unable to register gateway handler")
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *restPort),
		Handler: m,
	}

	return gwServer
}
```

##### using the default swagger file locations
```go
func initialiseSwaggerAPI(ctx context.Context, restPort, grpcPort, maxMessageSize *int) *http.Server {
	m := http.NewServeMux()

	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{},
	}))

	m.Handle("/", gwmux)
	m.Handle("/swagger-ui/", http.stripPrefix("/swagger-ui/", swaggerui.SwaggerUI()))
	m.Handle("/swagger.json", swaggerui.SwaggerFile())

	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(*maxMessageSize*1024*1024),
			grpc.MaxCallSendMsgSize(*maxMessageSize*1024*1024),
		),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer())),
		)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer())),
		)),
	}

	if rErr := my_gateway.RegisterServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("localhost:%d", *grpcPort), dialOpts); rErr != nil {
		log.WithError(rErr).Panic("unable to register gateway handler")
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *restPort),
		Handler: m,
	}

	return gwServer
}
```
