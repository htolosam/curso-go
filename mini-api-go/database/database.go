package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(dns string) error {
	var err error
	DB, err = sql.Open("mysql", dns)
	if err != nil {
		return fmt.Errorf("Error al conectar a la base de datos: %s", err)
	}
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Error al conectar a la base de datos: %s", err)
	}
	DB.SetMaxIdleConns(25) // configure maximum connections open
	DB.SetMaxOpenConns(10) // configure max inactive connections
	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
