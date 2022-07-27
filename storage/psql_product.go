package storage

import (
	"database/sql"
	"fmt"

	"github.com/jullentuss/go-db/pkg/product"
)

// buenas practicas en el uso de base de datos
const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT product_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct = `INSERT INTO 
		products(name, observations, price, created_at) 
		VALUES($1, $2, $3, $4) 
		RETURNING id`
	psqlGetAllProduct     = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	psqlGetAllProductByID = `SELECT id, name, observations, price, created_at, updated_at FROM products WHERE id = $1`
	psqlUpdateProduct     = `UPDATE products SET name = $1, observations = $2, price = $3, updated_at = $4 WHERE id = $5`
	psqlDeleteroduct      = `DELETE FROM products WHERE id = $1`
)

// psqlProduct usad for work with posgres - product
type psqlProduct struct {
	db *sql.DB
}

type scanner interface {
	Scan(dest ...any) error
}

// constructor return a new pointer of psqlProduct
func newPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}

// Migrate implemen the interface product.Storage
func (p *psqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
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

func (p *psqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)

	if err != nil {
		return err
	}

	fmt.Println("Se creo el producto correctamente")
	return nil
}

func (p *psqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
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

func (p *psqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteroduct)
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

func (p *psqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetAllProductByID)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

func (p *psqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
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
func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
