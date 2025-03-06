package application

import (
    "ModaVane/products/domain"
    "ModaVane/products/domain/ports"
)

type CreateProductUseCase struct {
    repo ports.ProductRepository
}

func NewCreateProductUseCase(repo ports.ProductRepository) *CreateProductUseCase {
    return &CreateProductUseCase{repo: repo}
}

func (uc *CreateProductUseCase) Execute(product domain.Product) (int, error) {
    return uc.repo.CreateProduct(product)
}
