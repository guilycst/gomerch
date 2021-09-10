package models

import (
	"gomerch/db"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Amount      int
}

func FindAll() []Product {
	gdb := db.GetConnection()
	var products []Product
	gdb.Find(&products)
	return products
}

func AddNewProduct(product Product) {
	gdb := db.GetConnection()
	gdb.Create(&product)
}

func UpdateProduct(product Product) {
	stored := FindProduct(product.ID)
	stored.Amount = product.Amount
	stored.Description = product.Description
	stored.Name = product.Name
	stored.Price = product.Price

	gdb := db.GetConnection()
	gdb.Save(&stored)
}

func FindProduct(ID uint) Product {
	gdb := db.GetConnection()
	stored := Product{}
	gdb.First(&stored, ID)
	return stored
}

func DeleteProduct(ID uint) {
	gdb := db.GetConnection()
	gdb.Delete(&Product{}, ID)
}
