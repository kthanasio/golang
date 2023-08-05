package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Category HAS MANY products
type Category struct {
	ID       int       `gorm:"primary-key"`
	Name     string    `db:"name" json:"name" gorm:"index:idx_name,unique"`
	Enabled  bool      `db:"enabled" json:"enabled"`
	Products []Product `gorm:"many2many:products_categories;"`
}

// Product HAS MANY categories
type Product struct {
	gorm.Model
	Name         string     `db:"name" json:"name"`
	Price        float64    `db:"price" json:"price"`
	Categories   []Category `gorm:"many2many:products_categories;"`
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

	// var category *Category
	// category = &Category{Name: "Moda Casa", Enabled: true}
	// err = db.Create(&category).Error
	// if err != nil {
	// 	panic("failed to Create Category")
	// }

	// var category2 *Category
	// category2 = &Category{Name: "Eletronicos", Enabled: true}
	// err = db.Create(&category2).Error
	// if err != nil {
	// 	panic("failed to Create Category")
	// }
	// var product *Product
	// product = &Product{Name: "Tapete", Price: 5500.99, Categories: []Category{*category, *category2}}
	// db.Create(&product)

	// var serialNumber *SerialNumber
	// serialNumber = &SerialNumber{Code: "##123A", ProductID: int(product.ID)}
	// err = db.Create(&serialNumber).Error
	// if err != nil {
	// 	panic(err)
	// }

	var categories []*Category
	db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories)
	for _, category := range categories {
		fmt.Printf("Category Name: %+v\n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("- [%v]: Product Name: %+v - Serial Number: %v\n", product.ID, product.Name, product.SerialNumber.Code)
		}
	}

	var products []*Product
	db.Model(&Product{}).Preload("SerialNumber").Preload("Categories").Find(&products)
	for _, product := range products {
		fmt.Printf("Product Name: %+v - Serial Number: %+v\n", product.Name, product.SerialNumber.Code)
		for _, category := range product.Categories {
			fmt.Printf("- [%v]: Category Name: %+v\n", category.ID, category.Name)
		}
	}
}
