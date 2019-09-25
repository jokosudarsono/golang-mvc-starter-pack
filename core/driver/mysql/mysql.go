package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct{}

var DB *sql.DB

func (msql *Mysql) Initialize() {
	connectionString := fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
