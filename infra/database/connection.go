package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const dataSource = "%s:%s@tcp(%s:3306)/%s?parseTime=true"

func Connect() *sql.DB {
	dataSourceName := fmt.Sprintf(dataSource, "dev", "dev", "mysqldb", "challenger")

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err.Error())
	}

	return db
}
