package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	connectionString := fmt.Sprintf("%s@tcp(%s:%s)/%s", user, host, port, dbname)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to MariaDB: ", connectionString)

	return db, err
}
