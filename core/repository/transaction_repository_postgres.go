package repository

import (
	"database/sql"
	"errors"
	"github.com/liverday/codebank/core/domain"
)

type TransactionsRepositoryPostgres struct {
	db *sql.DB
}

func NewTransactionRepositoryPostgres(db *sql.DB) *TransactionsRepositoryPostgres {
	return &TransactionsRepositoryPostgres{db: db}
}

func (t *TransactionsRepositoryPostgres) GetCreditCard(creditCard domain.CreditCard) (domain.CreditCard, error) {
	var c domain.CreditCard

	stmt, err := t.db.Prepare("select id, balance, balance_limit from credit_cards WHERE number = $1")
	if err != nil {
		return c, err
	}

	if err = stmt.QueryRow(creditCard.Number).Scan(&c.ID, &c.Balance, &c.Limit); err != nil {
		return c, errors.New("credit card does not exists")
	}

	return c, nil
}

func (t *TransactionsRepositoryPostgres) SaveTransaction(
	transaction domain.Transaction, 
	creditCard domain.CreditCard,
) error {
	stmt, err := t.db.Prepare(`
		insert into transactions (
			id, 
			credit_card_id, 
			amount, 
			status, 
			description, 
			store,
			created_at
		) values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7	
		)`)	

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardID,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt,
	);

	if err != nil {
		return err
	}

	if transaction.Status == "approved" {
		err = t.updateBalance(creditCard)
		if err != nil {
			return err
		}
	}

	err = stmt.Close()

	if err != nil {
		return err
	} 

	return nil
}

func (t *TransactionsRepositoryPostgres) updateBalance(
	creditCard domain.CreditCard,
) error {
	_, err := t.db.Exec("update credit_cards set balance = $1 WHERE id = $2", 
				creditCard.Balance, creditCard.ID)

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionsRepositoryPostgres) SaveCreditCard(creditCard domain.CreditCard) error {
	stmt, err := t.db.Prepare(`insert into credit_cards (
		id, 
		owner_name, 
		number,
		expiration_month, 
		expiration_year, 
		CVV, 
		balance, 
		balance_limit
	) values (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8
	)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.OwnerName,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit,
	)

	if err != nil {
		return err
	}
	
	err = stmt.Close()

	if err != nil {
		return err
	}

	return nil
}