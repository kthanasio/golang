package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Category ...
type Category struct {
	ID      int    `gorm:"primary-key"`
	Name    string `db:"name" json:"name" gorm:"index:idx_name,unique"`
	Enabled bool   `db:"enabled" json:"enabled"`
}

// Product BELONGS TO a category...
// Product HAS ONE SerialNumber
type Product struct {
	gorm.Model
	Name         string  `db:"name" json:"name"`
	Price        float64 `db:"price" json:"price"`
	Category     Category
	CategoryID   int
	SerialNumber SerialNumber
}

type SerialNumber struct {
	ID        int `gorm:"primary-key"`
	Code      string
	ProductID int
}

func main() {
	dsn := "root:admin@tcp(localhost:3306)/goexpert?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	if err != nil {
		panic("failed to migrate database")
	}

	var category *Category
	category = &Category{Name: "Eletronicos", Enabled: true}
	err = db.Create(&category).Error
	if err != nil {
		panic("failed to Create Category")
	}
	// var product *Product
	// product = &Product{Name: "Notebook", Price: 1000.00, CategoryID: category.ID}
	// db.Create(&product)

	// PRELOAD (análogo ao populate do mongodb)
	var product *Product
	product = &Product{Name: "TV", Price: 5000.00, CategoryID: category.ID}
	err = db.Create(&product).Error
	if err != nil {
		panic(err)
	}

	var serialNumber *SerialNumber
	serialNumber = &SerialNumber{Code: "123A", ProductID: int(product.ID)}
	err = db.Create(&serialNumber).Error
	if err != nil {
		panic(err)
	}

	var products []*Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Printf("Product Name: %+v - Category Name: %+v - SerialNumber: %+v\n", product.Name, product.Category.Name, product.SerialNumber.Code)
	}
}
