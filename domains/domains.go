package domains

import (
	"github.com/0xfzz/tuwit-backend/domains/auth"
	"github.com/0xfzz/tuwit-backend/domains/user"
	"github.com/0xfzz/tuwit-backend/utils/libfx"
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
