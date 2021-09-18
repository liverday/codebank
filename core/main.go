package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/liverday/codebank/core/infrastructure/grpc/server"
	"github.com/liverday/codebank/core/infrastructure/kafka"
	"github.com/liverday/codebank/core/infrastructure/repository"
	"github.com/liverday/codebank/core/usecase"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("can't load .env file")
	}
}

func main() {

	db := setupDb()
	kafkaProducer := setupKafkaProducer()
	useCase := setupTransactionUseCase(db, kafkaProducer)
	setupAndServeGrpc(useCase)

	defer db.Close()
}

func setupAndServeGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer(processTransactionUseCase)
	log.Println("Rodando gRPC")
	grpcServer.Serve()
}

func setupKafkaProducer() kafka.KafkaProducer {
	kafkaProducer := kafka.NewKafkaProducer()
	kafkaProducer.SetupProducer(
		os.Getenv("KafkaBootstrapServers"),
	)
	return kafkaProducer
}

func setupTransactionUseCase(db *sql.DB, kafkaProducer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionsRepository := repository.NewTransactionRepositoryPostgres(db)
	useCase := usecase.NewUseCaseTransaction(transactionsRepository, kafkaProducer)

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error connecting to database")
	}

	return db
}
