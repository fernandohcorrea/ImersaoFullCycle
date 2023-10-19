package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/codeedu/imersao/codepix-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewProduct(t *testing.T) {
	name := "Banco do Brasil"
	description := "Banco do Brasil"
	price := 1.99
	product, err := model.NewProduct(name, description, price)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(product.ID))
	require.Equal(t, product.Name, name)
	require.Equal(t, product.Description, description)

	_, err = model.NewProduct(name, "", price)
	require.NotNil(t, err)
}
