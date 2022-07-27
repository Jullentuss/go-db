package storage

import (
	"database/sql"
	"fmt"
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
)

// MySQLProduct usad for work with posgres - product
type MySQLProduct struct {
	db *sql.DB
}

// constructor return a new pointer of MySQLProduct
func NewMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

// Migrate implemen the interface product.Storage
func (p *MySQLProduct) Migrate() error {
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

// func (p *MySQLProduct) Create(m *product.Model) error {
// 	stmt, err := p.db.Prepare(mySQLMigrateProduct)
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	err = stmt.QueryRow(
// 		m.Name,
// 		stringToNull(m.Observations),
// 		m.Price,
// 		m.CreatedAt,
// 	).Scan(&m.ID)

// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Se creo el producto correctamente")
// 	return nil
// }

// func (p *MySQLProduct) GetAll() (product.Models, error) {
// 	stmt, err := p.db.Prepare(psqlGetAllProduct)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer stmt.Close()

// 	rows, err := stmt.Query()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	ms := make(product.Models, 0)

// 	for rows.Next() {
// 		m, err := scanRowProduct(rows)
// 		if err != nil {
// 			return nil, err
// 		}
// 		ms = append(ms, m)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return ms, nil
// }

// func (p *MySQLProduct) Delete(id uint) error {
// 	stmt, err := p.db.Prepare(psqlDeleteroduct)
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(id)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Se elimino el producto correctamente")
// 	return nil
// }

// func (p *MySQLProduct) GetByID(id uint) (*product.Model, error) {
// 	stmt, err := p.db.Prepare(psqlGetAllProductByID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer stmt.Close()

// 	return scanRowProduct(stmt.QueryRow(id))
// }

// func (p *MySQLProduct) Update(m *product.Model) error {
// 	stmt, err := p.db.Prepare(psqlUpdateProduct)
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	res, err := stmt.Exec(
// 		m.Name,
// 		stringToNull(m.Observations),
// 		m.Price,
// 		timeToNull(m.UpdatedAt),
// 		m.ID,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	rowsAffected, err := res.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	// alternativa para retornar error
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("no existe el producto con id: %d", m.ID)
// 	}

// 	fmt.Println("Se actualizo el producto correctamente")
// 	return nil
// }

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
