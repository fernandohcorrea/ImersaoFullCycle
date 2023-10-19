package usecase

import "github.com/codeedu/imersao/codepix-go/domain/model"

type ProductUseCase struct {
	ProductRepository model.ProductInterface
}

func (p *ProductUseCase) RegisterProduct(name string, description string, price float64) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	product_db, err := p.ProductRepository.RegisterProduct(product)
	if err != nil {
		return nil, err
	}

	return product_db, nil
}

func (p *ProductUseCase) FindProductById(id string) (*model.Product, error) {
	product, err := p.ProductRepository.FindProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
