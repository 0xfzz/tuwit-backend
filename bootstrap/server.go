package bootstrap

import (
	"context"
	"fmt"

	"github.com/0xfzz/tuwit-backend/config"
	"github.com/0xfzz/tuwit-backend/ent/migrate"
	"github.com/0xfzz/tuwit-backend/lib"
)

func StartServer(router lib.Router, database lib.Database) {
	ctx := context.Background()
	err := database.DB.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		panic("Error Auto Migrate")
	}
	addr := fmt.Sprintf(":%d", config.Get().APP.Port)
	router.Gin.Run(addr)
}
