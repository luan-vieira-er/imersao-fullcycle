package main

import (
	"codebank/domain"
	"codebank/infrastructure/repository"
	"codebank/usecase"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db := setupDB()
	defer db.Close()
	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Luan"
	cc.ExpirationYear = 2024
	cc.ExpirationMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDB(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDB(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	return useCase
}

func setupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"root",
		"codebank")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error connection to DB")
	}
	return db
}
