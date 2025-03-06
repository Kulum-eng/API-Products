package application

import (
    "ModaVane/products/domain/ports"
)

type DeleteProductUseCase struct {
    repo ports.ProductRepository
}

func NewDeleteProductUseCase(repo ports.ProductRepository) *DeleteProductUseCase {
    return &DeleteProductUseCase{repo: repo}
}

func (uc *DeleteProductUseCase) Execute(id int) error {
    return uc.repo.DeleteProduct(id)
}
