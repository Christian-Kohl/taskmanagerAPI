package connector

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DBInterface interface {
	Connect()
	Query()
	Statement()
}

var d *sql.DB

func main() {
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "testuser:TestUser1!@tcp(127.0.0.1:3306)/tasks")
	if err != nil {
		panic(err.Error())
	}

	return db
}
