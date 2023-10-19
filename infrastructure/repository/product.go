package repository

import (
	"fmt"

	"github.com/codeedu/imersao/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (r ProductRepositoryDb) RegisterProduct(product *model.Product) (*model.Product, error) {
	err := r.Db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductRepositoryDb) FindProductById(id string) (*model.Product, error) {
	var product model.Product

	r.Db.First(&product, "id = ?", id)

	if product.ID == "" {
		return nil, fmt.Errorf("no product found")
	}

	return &product, nil
}
