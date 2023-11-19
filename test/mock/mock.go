package mock

import (
	"testing"

	"github.com/0xfzz/tuwitt/domains"
	"github.com/0xfzz/tuwitt/domains/auth"
	"github.com/0xfzz/tuwitt/domains/user"
	"github.com/0xfzz/tuwitt/ent"
	"github.com/0xfzz/tuwitt/ent/enttest"
	"github.com/0xfzz/tuwitt/ent/migrate"
	"github.com/0xfzz/tuwitt/lib"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

type Database struct {
	DB *ent.Client
}

func NewMockDatabase(t *testing.T) lib.Database {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	return lib.Database{
		DB: client,
	}
}

var libModules = fx.Options(
	fx.Provide(NewMockDatabase),
	fx.Provide(lib.NewRouter),
	fx.Provide(lib.NewApi),
)

var Modules = fx.Options(
	libModules,
	domains.Module,
)

func RouterAuthController(t *testing.T, db lib.Database) *gin.Engine {
	router := lib.NewRouter()
	api := lib.NewApi(router)
	auth.NewAuthController(api, auth.NewAuthService(user.NewUserRepository(db))).AddRoute()
	return router.Gin

}
