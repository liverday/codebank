package usecase

import (
	"time"
	"github.com/liverday/codebank/core/domain"
	"github.com/liverday/codebank/core/dto"
)

type UseCaseTransaction struct {
	transactionsRepository domain.TransactionsRepository
}

func NewUseCaseTransaction(transactionsRepository domain.TransactionsRepository) UseCaseTransaction {
	return UseCaseTransaction{transactionsRepository: transactionsRepository}
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