package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	RegisterProduct(product *Product) (*Product, error)
	FindProductById(id string) (*Product, error)
}

type Product struct {
	Base        `valid:"required"`
	Name        string  `json:"name" gorm:"type:varchar(150)" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(150)" valid:"notnull"`
	Price       float64 `json:"price" gorm:"type:float" valid:"float, notnull"`
}

func (product *Product) isValid() error {
	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return err
	}
	return nil
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
	product.ID = uuid.NewV4().String()
	product.CreatedAt = time.Now()

	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
