package usecase

import (
	"fmt"
	"lsport/model"
	"lsport/repository"
)

type ProductUseCase struct {
	// repository
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}

}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := pu.repository.CreateProduct(product)
	if err != nil {
		fmt.Println("Error no CreateProduct => ", err)
		return model.Product{}, err

	}

	product.ID = id

	return product, nil

}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(id)

	if err != nil {
		fmt.Println("Error no GetProductById => ", err)
		return nil, err

	}

	return product, nil

}
