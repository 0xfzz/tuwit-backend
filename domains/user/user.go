package user

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewUserRepository,
	NewUserService,
)
