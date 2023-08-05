package main

import (
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

// Category HAS MANY products
type Category struct {
	ID      int    `gorm:"primary-key"`
	Name    string `db:"name" json:"name" gorm:"index:idx_name,unique"`
	Enabled bool   `db:"enabled" json:"enabled"`
}

func main() {
	dsn := "root:admin@tcp(localhost:3306)/goexpert?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Category{})
	if err != nil {
		panic("failed to migrate database")
	}

	PrintCategories(db)

	// transaction begin
	tx := db.Begin()
	var category *Category
	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, 2).Error
	if err != nil {
		panic(err)
	}
	category.Name = category.Name + " updated " + strconv.Itoa(category.ID)
	err = tx.Save(&category).Error
	if err != nil {
		panic(err)
	}
	tx.Commit()
	// transaction end

	PrintCategories(db)

}

func PrintCategories(db *gorm.DB) {
	var categories []*Category
	db.Model(&Category{}).Find(&categories)
	for _, category := range categories {
		fmt.Printf("Category Name: %+v\n", category.Name)
	}
}
