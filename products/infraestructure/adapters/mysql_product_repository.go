package adapters

import (
	"database/sql"
	"errors"

	"ModaVane/products/domain"
)

type MySQLProductRepository struct {
	DB *sql.DB
}

func NewMySQLProductRepository(db *sql.DB) *MySQLProductRepository {
	return &MySQLProductRepository{DB: db}
}

func (repo *MySQLProductRepository) CreateProduct(product domain.Product) (int, error) {
	res, err := repo.DB.Exec(
		"INSERT INTO products (name, description, price, size, color, category, stock) VALUES (?, ?, ?, ?, ?, ?, ?)",
		product.Name, product.Description, product.Price, product.Size, product.Color, product.Category, product.Stock,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (repo *MySQLProductRepository) GetProductByID(id int) (*domain.Product, error) {
	var product domain.Product
	err := repo.DB.QueryRow(
		"SELECT id, name, description, price, size, color, category, stock FROM products WHERE id = ?",
		id,
	).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Size, &product.Color, &product.Category, &product.Stock)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (repo *MySQLProductRepository) GetAllProducts() ([]domain.Product, error) {
	rows, err := repo.DB.Query("SELECT id, name, description, price, size, color, category, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Size, &p.Color, &p.Category, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *MySQLProductRepository) UpdateProduct(product domain.Product) error {
	_, err := repo.DB.Exec(
		"UPDATE products SET name=?, description=?, price=?, size=?, color=?, category=?, stock=? WHERE id=?",
		product.Name, product.Description, product.Price, product.Size, product.Color, product.Category, product.Stock, product.ID,
	)
	return err
}

func (repo *MySQLProductRepository) DeleteProduct(id int) error {
	res, err := repo.DB.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no se eliminó ningún registro")
	}

	return nil
}
