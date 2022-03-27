package repo

import (
	"fmt"
	"time"

	entities "github.com/gamze.sakallioglu/learningGo/homework-3-week-4-gamzesakallioglu/domain/entities"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

// returns a new AuthorRepository object with the given gorm.DB
func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

// Migrations keeps db schema up to date
func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&entities.Author{})

}

// Checks if a data exists with the same ID, if not inserts it to the database
func (a *AuthorRepository) InsertOneData(author entities.Author) error {
	// error handling
	err := a.db.Where(entities.Author{ID: author.ID}).Attrs(entities.Author{ID: author.ID, Name: author.Name}).FirstOrCreate(&author).Error
	if err != nil {
		return err
	}
	//
	a.db.Where(entities.Author{ID: author.ID}).Attrs(entities.Author{ID: author.ID, Name: author.Name}).FirstOrCreate(&author)
	return nil
}

func (a *AuthorRepository) InsertDatas(authors entities.Authors) {
	for _, v := range authors {
		a.InsertOneData(v)
	}
}

// Returns the author with the given ID
func (a *AuthorRepository) GetById(id int) (*entities.Author, error) {
	var author entities.Author
	// error handling
	err := a.db.Where(&entities.Author{ID: id}).First(&author).Error
	if err != nil {
		return nil, err
	}
	//
	a.db.Where(&entities.Author{ID: id}).First(&author)
	return &author, nil
}

// Get Author with books they published
func (a *AuthorRepository) GetAuthorsWithBook() (*entities.Authors, error) {
	var authors entities.Authors
	// error handling
	err := a.db.Preload("Books").Find(&authors).Error
	if err != nil {
		return nil, err
	}
	//
	a.db.Preload("Books").Find(&authors)
	return &authors, nil
}

// soft delete - change deleted_at as now rather than actual delete
// do not delete an author if they have books in the books table
func (a *AuthorRepository) DeleteById(authorId int) error {
	bookRepository := NewBookRepository(a.db)
	// error handling
	authors, err := bookRepository.getBooksByAuthorId(authorId)
	if err != nil {
		return err
	}
	//
	if len(*authors) > 0 {
		fmt.Println("This authors has books, cannot be deleted")
	} else {
		a.db.Model(&entities.Author{}).Where("Id = ?", authorId).Update("deleted_at", time.Now())
	}
	return nil
}
