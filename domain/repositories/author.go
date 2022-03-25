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

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&entities.Author{})

}

func (a *AuthorRepository) InsertOneData(author entities.Author) {
	a.db.Where(entities.Author{Id: author.Id}).Attrs(entities.Author{Id: author.Id, Name: author.Name}).FirstOrCreate(&author)
}

func (a *AuthorRepository) InsertDatas(authors entities.Authors) {
	for _, v := range authors {
		a.InsertOneData(v)
	}
}

func (a *AuthorRepository) GetById(id int) entities.Author {
	var author entities.Author
	a.db.Where(&entities.Author{Id: id}).First(&author)
	return author
}

func (a *AuthorRepository) GetAuthorWithBook(book entities.Book) entities.Author {
	var authorId = book.AuthorId
	var author entities.Author
	a.db.Where(&entities.Author{Id: authorId}).First(&author)
	return author
}

// soft delete - change deleted_at as now rather than actual delete
// do not delete an authors if they have books in the books table
func (a *AuthorRepository) DeleteById(authorId int) {
	bookRepository := NewBookRepository(a.db)
	authors := bookRepository.GetBooksWithAuthor(entities.Author{Id: authorId})
	if len(authors) > 0 {
		fmt.Println("This authors has books, cannot be deleted")
	} else {
		a.db.Model(&entities.Author{}).Where("Id = ?", authorId).Update("deleted_at", time.Now())
	}
}
