package applicacion

import (
    "ModaVane/products/domain"
    "ModaVane/products/domain/ports"
)

type GetProductUseCase struct {
    repo ports.ProductRepository
}

func NewGetProductUseCase(repo ports.ProductRepository) *GetProductUseCase {
    return &GetProductUseCase{repo: repo}
}

func (uc *GetProductUseCase) ExecuteByID(id int) (*domain.Product, error) {
    return uc.repo.GetProductByID(id)
}

func (uc *GetProductUseCase) ExecuteAll() ([]domain.Product, error) {
    return uc.repo.GetAllProducts()
}
