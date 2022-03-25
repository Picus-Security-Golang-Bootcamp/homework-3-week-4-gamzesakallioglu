package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id         int    `gorm:"primaryKey,autoIncrement"`
	Name       string `gorm:"varchar(150)"`
	TotalPage  int
	TotalStock int
	Price      float32
	StockCode  string `gorm:"varchar(50)"`
	ISBN       string `gorm:"varchar(50)"`
	Author     Author `gorm:"foreignKey:AuthorId"`
	AuthorId   int
}

type Books []Book

// Adding Hooks
func (b *Book) ToString() string {
	return fmt.Sprintf("ID: %v\nNAME: %s\nPAGE NUMBER: %v\nTOTAL STOCK: %v\nPRICE: %v\nSTOCK CODE: %s\nISBN: %s\n",
		b.Id, b.Name, b.TotalPage, b.TotalStock, b.Price, b.StockCode, b.ISBN)
}

func (b *Book) BeforeDelete(db *gorm.DB) (err error) {
	fmt.Printf("%s is being deleted", b.Name)
	return nil
}
