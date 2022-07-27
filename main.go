package main

import (
	"fmt"
	"log"

	"github.com/jullentuss/go-db/pkg/product"
	"github.com/jullentuss/go-db/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	serviceProduct := product.NewService(myStorage)
	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	fmt.Println(ms)

}
