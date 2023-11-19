package bootstrap

import (
	"github.com/0xfzz/tuwit-backend/domains"
	"github.com/0xfzz/tuwit-backend/lib"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	lib.Module,
	domains.Module,
)

func Start() {
	fx.New(
		Modules,
		fx.Invoke(
			StartServer,
		),
	)
}
