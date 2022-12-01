package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Conectar() (*sql.DB, error) {
	stringConexao := "root:root@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("Connected Database with sucess: %v", db)
	return db, nil
}
