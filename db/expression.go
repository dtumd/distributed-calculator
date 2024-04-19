package db

import (
	"context"
	"database/sql"
	"fmt"
	mdl "yc/distr-calc/model"

	"github.com/google/uuid"
)

func CreateExpressionsTable(ctx context.Context, db *sql.DB) error {

	const expressionsTable = `
	CREATE TABLE IF NOT EXISTS expressions(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid TEXT NOT NULL,
		value TEXT NOT NULL,
		status TEXT NOT NULL,
		result TEXT NOT NULL,
		user_id INTEGER NOT NULL,

		FOREIGN KEY(user_id) REFERENCES users (id)
	);`

	if _, err := db.ExecContext(ctx, expressionsTable); err != nil {
		return err
	}

	return nil
}

func InsertExpression(ctx context.Context, tx *sql.Tx, expr *mdl.Expression) (int64, error) {
	var q = `
	INSERT INTO expressions (uuid, value, status, result, user_id) values ($1, $2, $3, $4, $5)
	`

	result, err := tx.ExecContext(ctx, q, expr.Uuid, expr.Value, expr.Status, expr.Result, expr.UserID)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func SelectExpressions(ctx context.Context, db *sql.DB) ([]mdl.Expression, error) {
	var expressions []mdl.Expression
	var q = "SELECT id, uuid, value, status, result, user_id FROM expressions"

	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	id := 0

	for rows.Next() {
		e := mdl.Expression{}
		err := rows.Scan(&id, &e.Uuid, &e.Value, &e.Status, &e.Result, &e.UserID)
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, e)
	}

	return expressions, nil
}

func UpdateExpression(expr mdl.Expression) {
	//save(expr)

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

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	var q = `
	UPDATE expressions SET status=$1, result=$2 WHERE uuid=$3
	`
	_, err = tx.ExecContext(ctx, q, expr.Status, expr.Result, expr.Uuid)
	if err != nil {
		//	return err
		fmt.Print(err)
	}

	tx.Commit()

	//return nil
}

func GetExpressions() map[string][]mdl.Expression {
	//es := maps.Values(exprs)

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

	exps, err := SelectExpressions(ctx, db)
	if err != nil {
		//	return err
		fmt.Print(err)
	}

	r := map[string][]mdl.Expression{
		"Expressions": exps,
	}

	return r
}

func SaveExpression(value string, status string, result string, login string) mdl.Expression {
	fmt.Println("SaveExpression, user login: " + login)

	id := uuid.New().String()

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

	user, err := SelectUser(ctx, db, login)
	if err != nil {
		panic(err)
	}

	expression := mdl.Expression{Uuid: id, Status: status, Value: value, Result: result, UserID: user.ID}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	expressionID, err := InsertExpression(ctx, tx, &expression)
	if err != nil {
		panic(err)
	}

	//expression.ID = expressionID
	fmt.Println(expressionID)

	tx.Commit()

	return expression
}
