package storage

import (
	"database/sql"
	"fmt"

	"github.com/jullentuss/go-db/pkg/invoice"
	"github.com/jullentuss/go-db/pkg/invoiceheader"
	"github.com/jullentuss/go-db/pkg/invoiceitem"
)

type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice {
	return &MySQLInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

func (p *MySQLInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Printf("Factura Creada con Id %d\n", m.Header.ID)

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Printf("%d items creados\n", len(m.Items))

	return tx.Commit()
}
