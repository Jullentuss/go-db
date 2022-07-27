package storage

import (
	"database/sql"
	"fmt"

	"github.com/jullentuss/go-db/pkg/product"
)

// buenas practicas en el uso de base de datos
const (
	mySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	mySQLCreateProduct = `INSERT INTO 
		products(name, observations, price, created_at) 
		VALUES(?, ?, ?, ?)`
	mySQLGetAllProduct     = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	mySQLGetAllProductByID = `SELECT id, name, observations, price, created_at, updated_at FROM products WHERE id = ?`
	mySQLUpdateProduct     = `UPDATE products SET name = ?, observations = ?, price = ?, updated_at = ? WHERE id = ?`
	mySQLDeleteroduct      = `DELETE FROM products WHERE id = ?`
)

// mySQLProduct usad for work with posgres - product
type mySQLProduct struct {
	db *sql.DB
}

// constructor return a new pointer of mySQLProduct
func newMySQLProduct(db *sql.DB) *mySQLProduct {
	return &mySQLProduct{db}
}

// Migrate implemen the interface product.Storage
func (p *mySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de product ejecutada correctamente")
	return nil
}

func (p *mySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLCreateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	m.ID = uint(id)

	fmt.Printf("Se creo el producto correctamente con id: %d\n", m.ID)
	return nil
}

func (p *mySQLProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(mySQLGetAllProduct)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ms := make(product.Models, 0)

	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

func (p *mySQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(mySQLDeleteroduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Se elimino el producto correctamente")
	return nil
}

func (p *mySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(mySQLGetAllProductByID)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

func (p *mySQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLUpdateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// alternativa para retornar error
	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con id: %d", m.ID)
	}

	fmt.Println("Se actualizo el producto correctamente")
	return nil
}

// funcion helper
// func scanRowProduct(s scanner) (*product.Model, error) {
// 	m := &product.Model{}
// 	observationNull := sql.NullString{}
// 	updatedAtNull := sql.NullTime{}

// 	err := s.Scan(
// 		&m.ID,
// 		&m.Name,
// 		&observationNull,
// 		&m.Price,
// 		&m.CreatedAt,
// 		&updatedAtNull,
// 	)
// 	if err != nil {
// 		return &product.Model{}, err
// 	}

// 	m.Observations = observationNull.String
// 	m.UpdatedAt = updatedAtNull.Time

// 	return m, nil
// }
