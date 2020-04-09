//+build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/knwoop/swagger-sample/handler"
)

func InitializeApp() (appHandler *handler.AppHandler) {
	wire.Build(handler.WireSet)
	return
}
