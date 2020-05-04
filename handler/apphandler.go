package handler

import (
	"github.com/knwoop/swagger-sample/mock/restapi/operations"
)

type Handler struct {
	Hello *HelloHandler
}

func (h *Handler) Register(api *operations.GreeterAPI) {
	api.GetGreetingHandler = operations.GetGreetingHandlerFunc(h.Hello.hello)
}
