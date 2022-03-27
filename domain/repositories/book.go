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

// Migrations keeps db schema up to date
func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&entities.Book{})
}

// Checks db if a data exists with same ID, If not inserts it to database.
func (b *BookRepository) InsertOneData(book entities.Book) error {
	// error handling
	err := b.db.Where(entities.Book{ID: book.ID, StockCode: book.StockCode}, entities.Author{ID: book.AuthorId}).
		Attrs(entities.Book{ID: book.ID, Name: book.Name, TotalPage: book.TotalPage, TotalStock: book.TotalStock, Price: book.Price, StockCode: book.StockCode, ISBN: book.ISBN, AuthorId: book.AuthorId}).
		FirstOrCreate(&book).Error
	if err != nil {
		return err
	}
	//
	b.db.Where(entities.Book{ID: book.ID, StockCode: book.StockCode}, entities.Author{ID: book.AuthorId}).
		Attrs(entities.Book{ID: book.ID, Name: book.Name, TotalPage: book.TotalPage, TotalStock: book.TotalStock, Price: book.Price, StockCode: book.StockCode, ISBN: book.ISBN, AuthorId: book.AuthorId}).
		FirstOrCreate(&book)
	return nil
}

// InsertData inserts the books one by one in the books slice
func (b *BookRepository) InsertDatas(books entities.Books) {
	for _, v := range books {
		b.InsertOneData(v)
	}
}

// GetById returns the book with the id that passed
func (b *BookRepository) GetById(id int) (*entities.Book, error) {
	var book entities.Book
	// error handling
	err := b.db.Where(&entities.Book{ID: id}).Preload("Author").First(&book).Error
	if err != nil {
		return nil, err
	}
	//
	b.db.Where(&entities.Book{ID: id}).Preload("Author").First(&book)
	return &book, nil
}

// Returns the book slice with the info of their authors
func (b *BookRepository) GetBooksWithAuthor() (*entities.Books, error) {
	var books entities.Books
	// error handling
	err := b.db.Preload("Author").Find(&books).Error
	if err != nil {
		return nil, err
	}
	//
	b.db.Preload("Author").Find(&books)
	return &books, nil
}

// Returns the book slice that has the keyword
func (b *BookRepository) FindByName(name string) (entities.Books, error) {
	name = "%" + name + "%"
	var books entities.Books
	// error handling
	err := b.db.Where("name LIKE ?", name).Find(&books).Error
	if err != nil {
		return nil, err
	}
	//
	b.db.Where("name LIKE ?", name).Find(&books)
	return books, nil
}

// Returns the book slice that starts with the keyword
func (b *BookRepository) FindStartsWithName(name string) (*entities.Books, error) {
	name = "%" + name
	var books entities.Books
	// error handling
	err := b.db.Where("name LIKE ?", name).Find(&books).Error
	if err != nil {
		return nil, err
	}
	//
	b.db.Where("name LIKE ?", name).Find(&books)
	return &books, nil
}

// Returns the book slice that ends with the keyword
func (b *BookRepository) FindEndsWithName(name string) (*entities.Books, error) {
	name = name + "%"
	var books entities.Books
	// error handling
	err := b.db.Where("name LIKE ?", name).Find(&books).Error
	if err != nil {
		return nil, err
	}
	//
	b.db.Where("name LIKE ?", name).Find(&books)
	return &books, nil
}

// Returns the book slice that published by given author id
// Starts with lower case, only be used in this package (for checking if the author can be deleted or not)
func (b *BookRepository) getBooksByAuthorId(id int) (*entities.Books, error) {
	var books entities.Books
	// error handling
	err := b.db.Where(&entities.Book{AuthorId: id}).Find(&books).Error
	if err != nil {
		return nil, err
	}
	//
	b.db.Where(&entities.Book{AuthorId: id}).Find(&books)
	return &books, err
}

// Buy the book with given ID with given amount
func (b *BookRepository) BuyBook(id int, quantity int) error {
	book, _ := b.GetById(id)
	// old total stock of the book
	var bookStock = book.TotalStock
	if bookStock == 0 {
		fmt.Println("Sorry, This book is out of stock.\nYou can look for other books or come back later")
	} else if quantity > bookStock {
		err := b.updateStock(id, 0)
		if err != nil {
			return err
		}
		fmt.Println("We have only ", bookStock, " amount of ", book.Name, " books, so you bougth this much")
	} else if quantity <= bookStock {
		err := b.updateStock(id, (bookStock - quantity))
		if err != nil {
			return err
		}
		fmt.Println("You bougth ", quantity, " amounts of ", book.Name, " book")
	}
	return nil
}

// Updates the stock of the book data with given ID
func (b *BookRepository) updateStock(id, stock int) error {
	// error handling
	err := b.db.Model(&entities.Book{}).Where("Id = ?", id).Update("total_stock", stock).Error
	if err != nil {
		return err
	}
	//
	b.db.Model(&entities.Book{}).Where("Id = ?", id).Update("total_stock", stock)
	return nil
}

// soft delete - change deleted_at as now rather than actual delete
func (b *BookRepository) DeleteById(bookId int) error {
	// error handling
	err := b.db.Model(&entities.Book{}).Where("Id = ?", bookId).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}
	//
	b.db.Model(&entities.Book{}).Where("Id = ?", bookId).Update("deleted_at", time.Now())
	return nil
}
