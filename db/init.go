package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "dc.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	if err = CreateExpressionsTable(ctx, db); err != nil {
		panic(err)
	}

	if err = CreateUserTable(ctx, db); err != nil {
		panic(err)
	}

	InitSettings() // TODO
}
