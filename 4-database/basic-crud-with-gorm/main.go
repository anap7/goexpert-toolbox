package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model //vai inserir as colunas created_at, updated_at, deleted_at
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Mouse",
		Price:      1000.00,
		CategoryID: 1,
	})

	// select one
	var product Product
	db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	var products []Product

	fmt.Println("================SELECT COMUM====================")
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}

	fmt.Println("================SELECT COM PAGINACAO====================")
	db.Limit(2).Offset(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	fmt.Println("================SELECT COM WHERE====================")
	db.Where("price > ?", 90).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
	db.Where("name LIKE ?", "book").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	//update
	var p Product
	db.First(&p, 1)
	p.Name = "Notebook"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	//delete
	db.Delete(&p2)



}
