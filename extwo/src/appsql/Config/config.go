package config

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "168.63.242.2"
	port     = 6969
	user     = "user"
	password = "2019predator030317"
	database = "predator_4_beta"
)

func DBOp() (db *sql.DB){
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",server, user, password, port, database)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	return conn
}