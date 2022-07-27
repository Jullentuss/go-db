package main

import (
	"log"

	"github.com/jullentuss/go-db/pkg/product"
	"github.com/jullentuss/go-db/storage"
)

func main() {
	storage.NewMySQLDB()

	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("producto.Migrate: %v", err)
	}

	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	serviceHeader := product.NewService(storageHeader)

	if err := serviceHeader.Migrate(); err != nil {
		log.Fatalf("header.Migrate: %v", err)
	}

	storageItem := storage.NewMySQLInvoiceItem(storage.Pool())
	serviceItem := product.NewService(storageItem)

	if err := serviceItem.Migrate(); err != nil {
		log.Fatalf("item.Migrate: %v", err)
	}
}
