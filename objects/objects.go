package objects

import (
	"database/sql"
	"taskmanagerAPI/connector"
)

var db *sql.DB

func main() {
	db = connector.Connect()
}
