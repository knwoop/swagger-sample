package handler

import "github.com/knwoop/swagger-sample/mock/restapi/operations"

type AppHandler struct {
	hello *HelloHandler
}

func NewAppHandler(h *HelloHandler) *AppHandler {
	return &AppHandler{hello: h}
}

func (app *AppHandler) Register(api *operations.GreeterAPI) {
	api.GetGreetingHandler = operations.GetGreetingHandlerFunc(app.hello.hello)
}
