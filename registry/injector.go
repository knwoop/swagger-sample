//+build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/knwoop/swagger-sample/handler"
)

func InitializeApp() (appHandler *handler.Handler) {
	wire.Build(
		handler.WireSet,
		wire.Struct(new(handler.Handler), "*"),
	)
	return
}
