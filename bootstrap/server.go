package bootstrap

import (
	"context"
	"fmt"

	"github.com/0xfzz/tuwitt/lib"
)

func StartServer(router lib.Router, config lib.Config, database lib.Database) {
	ctx := context.Background()
	err := database.DB.Schema.Create(ctx)
	if err != nil {
		panic("Error Auto Migrate")
	}
	addr := fmt.Sprintf(":%d", config.APP.Port)
	router.App.Run(addr)
}
