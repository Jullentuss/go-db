package main

import (
	"log"

	"github.com/jullentuss/go-db/pkg/invoice"
	"github.com/jullentuss/go-db/pkg/invoiceheader"
	"github.com/jullentuss/go-db/pkg/invoiceitem"
	"github.com/jullentuss/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())

	storageInvoice := storage.NewPSqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Steven Clavijo",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 1},
			&invoiceitem.Model{ProductID: 2},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

	//fmt.Println(m)
}
