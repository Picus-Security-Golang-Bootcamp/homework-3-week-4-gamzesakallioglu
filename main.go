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

	bookRepo := repo.NewBookRepository(db)
	bookRepo.Migrations()
	bookRepo.InsertDatas(bookList)

	authorRepo := repo.NewAuthorRepository(db)
	authorRepo.Migrations()
	authorRepo.InsertDatas(authorList)

}
