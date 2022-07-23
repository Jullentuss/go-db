package storage

import (
	"database/sql"
	"fmt"
)

// buenas practicas en el uso de base de datos
const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now(),
		update_at TIMESTAMP,
		CONSTRAINT invoice_item_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
)

// PsqlInvoiceItem usad for work with posgres - product
type PsqlInvoiceItem struct {
	db *sql.DB
}

// constructor return a new pointer of PsqlInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Migrate implemen the interface invoice_item.Storage
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de InvoiceItem ejecutada correctamente")
	return nil
}
