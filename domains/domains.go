package domains

import (
	"github.com/0xfzz/tuwitt/domains/auth"
	"github.com/0xfzz/tuwitt/domains/user"
	"github.com/0xfzz/tuwitt/utils/libfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	user.Module,
	auth.Module,
	fx.Invoke(
		fx.Annotate(
			InitAllRoutes,
			fx.ParamTags(`group:"controllers"`),
		),
	),
)

func InitAllRoutes(controllers []libfx.Controller) {
	for _, controller := range controllers {
		controller.AddRoute()
	}
}
