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

	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
