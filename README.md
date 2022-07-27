
# Product Migrations 
```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)

if err := serviceProduct.Migrate(); err != nil {
	log.Fatalf("product.Migrate: %v", err)
}
```

# InvoiceHeaders Migrations
```go
storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

if err := serviceInvoiceHeader.Migrate(); err != nil {
	log.Fatalf("invoiceHeader.Migrate: %v", err)
}
```

# InvoiceItems Migrations
```go
storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

if err := serviceInvoiceItem.Migrate(); err != nil {
	log.Fatalf("invoiceItem.Migrate: %v", err)
}
```

# insertar datos de producto 
```go
m := &product.Model{
		Name:         "Curso de DB con GO",
		Price:        70,
		Observations: "Curso disponible",
	}

	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%+v\n", m)
```

# consultar 1 dato
```go
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(3)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay un producto con el id")
	case err != nil:
		log.Fatalf("product.GetByID: %v\n", err)
	default:
		fmt.Println(m)
	}
```

# consultar varios registros
```go
storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)
```


# actualizar un dato
```go
storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		ID:           3,
		Name:         "Curso Python",
		Observations: "Curso de python desde 0",
		Price:        50,
	}

	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
```


# eliminar un producto
```go
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(3)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
```


# implementacion transaccion postgresql
```go
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
	```


---

# MySQL Migrate
```go
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
```


# MySQL Create 

```go
	storage.NewMySQLDB()
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:  "Curso de GO desde 0",
		Price: 70,
	}

	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%+v\n", m)
```

# MySQL get by id
```go
	storage.NewMySQLDB()
	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)
```
