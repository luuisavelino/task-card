package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Connection interface {
	Connect() (*sql.DB, error)
}

type MySQLConnection struct {
	user     string
	password string
	host     string
	port     string
	database string
}

func (c *MySQLConnection) Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.user, c.password, c.host, c.port, c.database))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectsWithDatabase() *sql.DB {
	mysql := MySQLConnection{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		database: os.Getenv("DB_DATABASE"),
	}

	db, err := mysql.Connect()

	if err != nil {
		panic(err.Error())
	}

	return db
}
