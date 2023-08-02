package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	data := Product{
		Name:  "Camista",
		Price: 99.99,
	}
	product, err := NewProduct(data.Name, data.Price)
	assert.NotNil(t, product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.Equal(t, data.Name, product.Name)
	assert.NotEmpty(t, product.Price)
	assert.Equal(t, data.Price, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestNewProductWhenNameIsRequired(t *testing.T) {
	data := Product{
		Name:  "",
		Price: 99.99,
	}
	product, err := NewProduct(data.Name, data.Price)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Name is required")
	assert.Nil(t, product)
}

// func TestNewProductWhenIDIsRequired(t *testing.T) {
// 	data := Product{
// 		Name:  "Camiseta",
// 		Price: 100.0,
// 	}
// 	product, err := NewProduct(data.Name, data.Price)
// 	assert.NotNil(t, err)
// 	assert.Nil(t, product)
// }

func TestNewProductWhenPriceIsRequired(t *testing.T) {
	data := Product{
		Name:  "Camiseta",
		Price: 0,
	}
	product, err := NewProduct(data.Name, data.Price)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
	assert.Nil(t, product)
}

func TestNewProductWhenInvalidPrice(t *testing.T) {
	data := Product{
		Name:  "Camiseta",
		Price: -1,
	}
	product, err := NewProduct(data.Name, data.Price)
	assert.NotNil(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
	assert.Nil(t, product)
}

func TestNewProductValidate(t *testing.T) {
	data := Product{
		Name:  "Camiseta",
		Price: 99.9,
	}
	product, err := NewProduct(data.Name, data.Price)
	assert.NotNil(t, product)
	assert.Nil(t, err)
	assert.Nil(t, product.Validate())

}
