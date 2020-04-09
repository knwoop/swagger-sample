package handler

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/knwoop/swagger-sample/mock/restapi/operations"
)

type HelloHandler struct {
}

func (_ *HelloHandler) hello(params operations.GetGreetingParams) middleware.Responder {
	name := swag.StringValue(params.Name)
	if name == "" {
		name = "World"
	}

	greeting := fmt.Sprintf("Hello, %s!", name)
	return operations.NewGetGreetingOK().WithPayload(greeting)
}
