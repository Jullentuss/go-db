package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

type Driver string

func New(d Driver) {
	switch d {
	case MySQL:
		NewMySQLDB()
	case Postgres:
		NewPostgresDB()
	}
}

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

func NewMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/dbgo?parseTime=true")
		if err != nil {
			// detiene la ejecucion de la aplicacion
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			// detiene la ejecucion de la aplicacion
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Connected to mysql")
	})
}

// retorna unica instancia de db
func Pool() *sql.DB {
	return db
}

// helper para manipular los datos nulos en base de datos
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}
