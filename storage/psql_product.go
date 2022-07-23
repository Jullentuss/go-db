package storage

import (
	"database/sql"
	"fmt"
)

// buenas practicas en el uso de base de datos
const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observationes VARCHAR(100),
		price INT NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		CONSTRAINT product_id_pk PRIMARY KEY (id)
	)`
)

// PsqlProduct usad for work with posgres - product
type PsqlProduct struct {
	db *sql.DB
}

// constructor return a new pointer of PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implemen the interface product.Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de product ejecutada correctamente")
	return nil
}
