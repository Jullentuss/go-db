package storage

import (
	"database/sql"
	"fmt"
)

// buenas practicas en el uso de base de datos
const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS Invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		CONSTRAINT Invoice_headers_id_pk PRIMARY KEY (id)
	)`
)

// PsqlInvoiceHeader usad for work with posgres - product
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// constructor return a new pointer of PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate implemen the interface invoice_header.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de InvoiceHeader ejecutada correctamente")
	return nil
}
