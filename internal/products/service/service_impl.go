package service

import (
	"test-be-dbo/internal/models"
	"test-be-dbo/internal/products/repository"
)

type ProductService struct {
	productRepos repository.Repository
}

func NewProductService(productRepos repository.Repository) Service {
	return &ProductService{productRepos: productRepos}
}

func (p *ProductService) GetAll() ([]models.ProductResponse, error) {
	result, err := p.productRepos.GetAll()
	if err != nil {
		return nil, err
	}

	var response []models.ProductResponse
	for _, product := range result {
		productResponse := models.ProductResponse{
			ID:        product.ID.String(),
			Name:      product.Name,
			Qty:       product.Qty,
			Price:     product.Price,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}
		response = append(response, productResponse)
	}

	return response, nil
}
