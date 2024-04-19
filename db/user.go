package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	mdl "yc/distr-calc/model"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name string, login string, password string) error {
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "dc.db")
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	pwd, err := generate(password)
	if err != nil {
		return err
	}

	user := &mdl.User{
		Name:     name,
		Login:    login,
		Password: pwd,
	}

	userID, err := InsertUser(ctx, db, user)

	if err != nil {
		log.Println("user already exists")
		tx.Rollback()
		return err
	} else {
		user.ID = userID
	}

	tx.Commit()

	return nil
}

func CheckPassword(login string, password string) error {
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "dc.db")
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	userFromDB, err := SelectUser(ctx, db, login)
	if err != nil {
		return err
	}

	err = Compare(userFromDB.Password, password)
	if err != nil {
		log.Println("auth fail")
		return err
	}

	log.Println("auth success")

	return nil
}

func GetUserByLogin(login string) (mdl.User, error) {
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "dc.db")
	if err != nil {
		return mdl.User{}, err
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		return mdl.User{}, err
	}

	user, err := SelectUser(ctx, db, login)
	if err != nil {
		return mdl.User{}, err
	}

	return user, nil
}

func CreateUserTable(ctx context.Context, db *sql.DB) error {
	const usersTable = `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		login TEXT UNIQUE,
		password TEXT
	);`

	if _, err := db.ExecContext(ctx, usersTable); err != nil {
		return err
	}

	return nil
}

func InsertUser(ctx context.Context, db *sql.DB, user *mdl.User) (int64, error) {
	var q = `
	INSERT INTO users (name, login, password) values ($1, $2, $3)
	`
	result, err := db.ExecContext(ctx, q, user.Name, user.Login, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func SelectUser(ctx context.Context, db *sql.DB, login string) (mdl.User, error) {
	fmt.Println("SelectUser, user login: " + login)
	var (
		user mdl.User
		err  error
	)

	var q = "SELECT id, name, login, password FROM users WHERE login=$1"
	err = db.QueryRowContext(ctx, q, login).Scan(&user.ID, &user.Name, &user.Login, &user.Password)
	return user, err
}

func Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}
