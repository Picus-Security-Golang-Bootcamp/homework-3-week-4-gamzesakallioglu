package repo

import (
	"fmt"
	"time"

	entities "github.com/gamze.sakallioglu/learningGo/homework-3-week-4-gamzesakallioglu/domain/entities"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&entities.Book{})
}

func (b *BookRepository) InsertOneData(book entities.Book) {
	b.db.Where(entities.Book{Id: book.Id, StockCode: book.StockCode}).
		Attrs(entities.Book{Id: book.Id, Name: book.Name, TotalPage: book.TotalPage, TotalStock: book.TotalStock, Price: book.Price, StockCode: book.StockCode, ISBN: book.ISBN, AuthorId: book.AuthorId}).
		FirstOrCreate(&book)
}

func (b *BookRepository) InsertDatas(books entities.Books) {
	for _, v := range books {
		b.InsertOneData(v)
	}
}

func (b *BookRepository) GetById(id int) entities.Book {
	var book entities.Book
	b.db.Where(&entities.Book{Id: id}).First(&book)
	return book
}

func (b *BookRepository) FindByName(name string) entities.Books {
	name = "%" + name + "%"
	var books entities.Books
	b.db.Where("name LIKE ?", name).Find(&books)
	return books
}

func (b *BookRepository) FindStartsWithName(name string) entities.Books {
	name = "%" + name
	var books entities.Books
	b.db.Where("name LIKE ?", name).Find(&books)
	return books
}

func (b *BookRepository) FindEndsWithName(name string) entities.Books {
	name = name + "%"
	var books entities.Books
	b.db.Where("name LIKE ?", name).Order("id asc").Find(&books)
	return books
}

func (b *BookRepository) GetBooksWithAuthor(author entities.Author) entities.Books {
	var authorId = author.Id
	var books entities.Books
	b.db.Where(&entities.Book{AuthorId: authorId}).Order("id asc").Find(&books)
	return books
}

func (b *BookRepository) BuyBook(book entities.Book, quantity int) {
	bookDb := b.GetById(book.Id)
	var bookStock = bookDb.TotalStock
	if bookStock == 0 {
		fmt.Println("This book is out of stock")
	} else if quantity > bookStock {
		b.updateStock(book, 0)
		fmt.Println("We have only ", bookStock, " amount of ", bookDb.Name, " books, so you bougth this much")
	} else if quantity <= bookStock {
		b.updateStock(book, (bookStock - quantity))
		fmt.Println("You bougth ", quantity, " amounts of ", bookDb.Name, " book")
	}
}

func (b *BookRepository) updateStock(book entities.Book, stock int) {
	bookId := book.Id
	b.db.Model(&entities.Book{}).Where("Id = ?", bookId).Update("total_stock", stock)
}

// soft delete - change deleted_at as now rather than actual delete
func (b *BookRepository) DeleteById(bookId int) {
	b.db.Model(&entities.Book{}).Where("Id = ?", bookId).Update("deleted_at", time.Now())
}
