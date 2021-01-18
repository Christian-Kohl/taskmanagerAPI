package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"taskmanagerAPI/routing"
)

type Post struct {
	Task_id   string `json:"task_id"`
	Task_name string `json:"task_name"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "testuser:TestUser1!@tcp(127.0.0.1:3306)/tasks")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router.Route
}
