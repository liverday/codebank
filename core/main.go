package main

import (
	"database/sql"
	"fmt"
	"github.com/liverday/codebank/core/usecase"
	"github.com/liverday/codebank/core/infrastructure/repository"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb
	defer db.Close()
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionsRepository := repository.NewTransactionRepositoryPostgres(db)
	useCase := usecase.NewUseCaseTransaction(transactionsRepository)
	
	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprint("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error connecting to database")
	}

	return db
}