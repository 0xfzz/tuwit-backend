package bootstrap

import (
	"github.com/0xfzz/tuwitt/domains"
	"github.com/0xfzz/tuwitt/lib"
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
