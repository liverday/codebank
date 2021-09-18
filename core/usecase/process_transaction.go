package usecase

import (
	"encoding/json"
	"os"
	"time"

	"github.com/liverday/codebank/core/domain"
	"github.com/liverday/codebank/core/dto"
	"github.com/liverday/codebank/core/infrastructure/kafka"
)

type UseCaseTransaction struct {
	transactionsRepository domain.TransactionsRepository
	kafkaProducer          kafka.KafkaProducer
}

func NewUseCaseTransaction(transactionsRepository domain.TransactionsRepository, kafkaProducer kafka.KafkaProducer) UseCaseTransaction {
	return UseCaseTransaction{transactionsRepository: transactionsRepository, kafkaProducer: kafkaProducer}
}

func (u UseCaseTransaction) ProcessTransaction(newTransactionDto dto.NewTransaction) (domain.Transaction, error) {
	creditCard := u.hydrateCreditCard(newTransactionDto)
	ccBalanceAndLimit, err := u.transactionsRepository.GetCreditCard(*creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}

	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	t := u.newTransaction(newTransactionDto, ccBalanceAndLimit)

	t.ProcessAndValidate(creditCard)

	err = u.transactionsRepository.SaveTransaction(*t, *creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}

	newTransactionDto.ID = t.ID
	newTransactionDto.CreatedAt = t.CreatedAt

	newTransactionJson, err := json.Marshal(newTransactionDto)

	if err != nil {
		return domain.Transaction{}, err
	}

	err = u.kafkaProducer.Publish(string(newTransactionJson), os.Getenv("KafkaTransactionsTopic"))

	if err != nil {
		return domain.Transaction{}, err
	}

	return *t, nil
}

func (u UseCaseTransaction) hydrateCreditCard(newTransactionDto dto.NewTransaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.OwnerName = newTransactionDto.OwnerName
	creditCard.Number = newTransactionDto.Number
	creditCard.ExpirationMonth = newTransactionDto.ExpirationMonth
	creditCard.ExpirationYear = newTransactionDto.ExpirationYear
	creditCard.CVV = newTransactionDto.CVV
	return creditCard
}

func (u UseCaseTransaction) newTransaction(newTransactionDto dto.NewTransaction, creditCard domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.CreditCardID = creditCard.ID
	t.Amount = newTransactionDto.Amount
	t.Store = newTransactionDto.Store
	t.Description = newTransactionDto.Description
	t.CreatedAt = time.Now()
	return t
}
