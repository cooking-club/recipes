package db

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var db *bun.DB

func InitDB() {
	ctx := context.Background()

	sqldb, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/recipes?parseTime=true&charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	defer sqldb.Close()

	// create db instance
	db = bun.NewDB(sqldb, mysqldialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	models := []any{
		(*Department)(nil),
		(*Room)(nil),
		(*Professor)(nil),
		(*GroupRecord)(nil),
		(*Group)(nil),
		(*Record)(nil),
	}

	for _, m := range models {
		_, err = db.NewCreateTable().Model(m).IfNotExists().Exec(ctx)
		if err != nil {
			panic(err)
		}
	}
}
