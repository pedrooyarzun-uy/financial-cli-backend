package db

import (
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var once sync.Once

func Init() {
	once.Do(func() {
		dsn := os.Getenv("DB_HOST")

		var err error
		DB, err = sqlx.Open("mysql", dsn)

		if err != nil {
			log.Fatal("String connection failed: ", err)
		}

		if err = DB.Ping(); err != nil {
			log.Fatal("Ping didn't work", err)
		}
	})
}

func Close() {
	if err := DB.Close(); err != nil {
		log.Fatal("Close failed", err)
	}
}
