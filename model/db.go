package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


const (
	host = "localhost"
	port = 5430
	user = "decagon"
	dbname = "tech"

)

// Db Inject database instance into model layer for easy testing
var Db *sql.DB

func init()  {
	var err error
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s dbname = %s sslmode = disable", host,port,user,dbname)
	Db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return
	}
	fmt.Println("database connected")
}