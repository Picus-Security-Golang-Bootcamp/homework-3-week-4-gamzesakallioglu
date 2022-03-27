package main

import (
	"fmt"
	"log"

	postgres "github.com/gamze.sakallioglu/learningGo/homework-3-week-4-gamzesakallioglu/common/db"
	csv_operations "github.com/gamze.sakallioglu/learningGo/homework-3-week-4-gamzesakallioglu/csv"
	repo "github.com/gamze.sakallioglu/learningGo/homework-3-week-4-gamzesakallioglu/domain/repositories"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	file_book := "books.csv"
	file_author := "authors.csv"
	bookList, err := csv_operations.ReadBooksCsv(file_book)
	if err != nil {
		fmt.Println(err)
	}

	authorList, err := csv_operations.ReadAuthorsCsv(file_author)
	if err != nil {
		fmt.Println(err)
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("cannot connect to postgres ", err)
	}

	authorRepo := repo.NewAuthorRepository(db)
	authorRepo.Migrations()
	authorRepo.InsertDatas(authorList)

	bookRepo := repo.NewBookRepository(db)
	bookRepo.Migrations()
	bookRepo.InsertDatas(bookList)

	books, err := bookRepo.GetBooksWithAuthor() // Books will be printed with the author
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range *books {
		fmt.Println(v.ToString())
	}

	authors, err := authorRepo.GetAuthorsWithBook() // Authors will be printed with their books
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range *authors {
		fmt.Println(v.ToString())
	}

}
