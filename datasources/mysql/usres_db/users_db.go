package usres_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_database_host     = "mysql_database_host"
	mysql_database_port     = "mysql_database_port"
	mysql_database_user     = "mysql_database_user"
	mysql_database_password = "mysql_database_password"
	mysql_database          = "mysql_database"
)

var (
	Client            *sql.DB
	database_host     = os.Getenv(mysql_database_host)
	database_port     = os.Getenv(mysql_database_port)
	database_user     = os.Getenv(mysql_database_user)
	database_password = os.Getenv(mysql_database_password)
	database          = os.Getenv(mysql_database)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		database_user,
		database_password,
		database_host,
		database_port,
		database,
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured")
}
