package main

import (
	"database/sql"
	"fmt"

	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct() *Product {

	product := &Product{}
	err := faker.FakeData(&product)
	if err != nil {
		panic(err)
	}
	return product
}

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct()
	insertProduct(db, product)
	product.Name = product.Name + " updated"
	updateProduct(db, product)

	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: id[%v] - Price: [%.2f] - Name: [%v]\n", p.ID, p.Price, p.Name)

	products, err := selectProducts(db)
	if err != nil {
		panic(err)
	}

	for _, p := range *products {
		fmt.Printf("Name: [%v] - Price: %.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) {
	stmt, err := db.Prepare("insert into products(id,name,price) values (?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		panic(err)
	}
}

func updateProduct(db *sql.DB, product *Product) {
	stmt, err := db.Prepare("update products set name = ? ,price = ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		panic(err)
	}
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil

}

func selectProducts(db *sql.DB) (*[]Product, error) {
	rows, err := db.Query("select id, name, price from products order by price desc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return &products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil

}
