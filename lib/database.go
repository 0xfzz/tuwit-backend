package lib

import (
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/0xfzz/tuwit-backend/config"
	"github.com/0xfzz/tuwit-backend/ent"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	DB *ent.Client
}

func NewDatabase() Database {
	config := config.Get()
	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)
	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		panic(fmt.Sprintf("Error while establishing database connection : %s", err))
	}
	return Database{
		DB: client,
	}

}
