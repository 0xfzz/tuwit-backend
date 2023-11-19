package auth

import (
	"github.com/0xfzz/tuwitt/utils/libfx"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewAuthService,
	fx.Annotate(
		NewAuthController,
		fx.As(new(libfx.Controller)),
		fx.ResultTags(`group:"controllers"`),
	),
)
