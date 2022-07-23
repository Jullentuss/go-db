package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// solo se ejecuta una vez, singleton
func NewPostgresDB() {
	once.Do(func() {
		connStr := "postgres://skoll:password@localhost/dbgo?sslmode=disable"
		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			// detiene la ejecucion de la aplicacion
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			// detiene la ejecucion de la aplicacion
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Connected to postgresql")
	})
}

// retorna unica instancia de db
func Pool() *sql.DB {
	return db
}
