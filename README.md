
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
serviceInvoiceHeader := product.NewService(storageInvoiceHeader)

if err := serviceInvoiceHeader.Migrate(); err != nil {
	log.Fatalf("invoiceHeader.Migrate: %v", err)
}
```

# InvoiceItems Migrations
```go
storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := product.NewService(storageInvoiceItem)

if err := serviceInvoiceItem.Migrate(); err != nil {
	log.Fatalf("invoiceItem.Migrate: %v", err)
}
```