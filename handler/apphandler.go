package handler

import "github.com/knwoop/swagger-sample/mock/restapi/operations"

type AppHandler struct {
	HelloHandler
	api operations.GreeterAPI
}

func (a *AppHandler)Handler(api *operations.GreeterAPI) {
	api.GetGreetingHandler = operations.GetGreetingHandlerFunc(a.HelloHandler.hello)
}
