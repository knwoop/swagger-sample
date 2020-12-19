package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/justinas/alice"
	"github.com/knwoop/swagger-sample/handler/middlewares"
	"github.com/knwoop/swagger-sample/mock/restapi"
	"github.com/knwoop/swagger-sample/mock/restapi/operations"
	"github.com/knwoop/swagger-sample/registry"
)

var portFlag = flag.Int("port", 3000, "port to run this serivice on")

func main() {
	// load embedded swagger file

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewGreeterAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	registry.InitializeApp().Register(api)

	handler := alice.New(
		middlewares.Middleware1,
		middlewares.Middleware2,
		middlewares.Middleware2,
	).Then(api.Serve(nil))

	server.SetHandler(handler)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
