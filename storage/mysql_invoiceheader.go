package storage

import (
	"database/sql"
	"fmt"

	"github.com/jullentuss/go-db/pkg/invoiceheader"
)

// buenas practicas en el uso de base de datos
const (
	mySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP
	)`
	mySQLCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES(?)`
)

// MySQLInvoiceHeader usad for work with posgres - product
type MySQLInvoiceHeader struct {
	db *sql.DB
}

// constructor return a new pointer of MySQLInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

// Migrate implemen the interface invoice_header.Storage
func (p *MySQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceHeader)
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

func (p *MySQLInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceHeader)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(m.Client)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	m.ID = uint(id)
	return nil
}
