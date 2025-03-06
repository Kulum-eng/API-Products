package application

import (
    "ModaVane/products/domain"
    "ModaVane/products/domain/ports"
)

type UpdateProductUseCase struct {
    repo ports.ProductRepository
}

func NewUpdateProductUseCase(repo ports.ProductRepository) *UpdateProductUseCase {
    return &UpdateProductUseCase{repo: repo}
}

func (uc *UpdateProductUseCase) Execute(product domain.Product) error {
    return uc.repo.UpdateProduct(product)
}
