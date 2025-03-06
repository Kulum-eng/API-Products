package ports

import "ModaVane/products/domain"

type ProductRepository interface {
    CreateProduct(product domain.Product) (int, error)
    GetProductByID(id int) (*domain.Product, error)
    GetAllProducts() ([]domain.Product, error)
    UpdateProduct(product domain.Product) error
    DeleteProduct(id int) error
}
