package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ProductGORM struct {
	gorm.Model
	Code  string
	Price float64
}

func main() {
	dsn := "root:admin@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&ProductGORM{})

	// Create
	db.Create(&ProductGORM{Code: "D42", Price: 100})

	// CreateBatch
	db.Create([]*ProductGORM{
		{Code: "A", Price: 10.0},
		{Code: "B", Price: 100.0},
		{Code: "C", Price: 1000.0},
		{Code: "D", Price: 100.9},
	})

	// Read
	var product ProductGORM
	db.First(&product, 1) // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(ProductGORM{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 201, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
