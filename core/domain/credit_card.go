package domain

import (
	"time"
	uuid "github.com/satori/go.uuid"
)


type CreditCard struct {
	ID string
	OwnerName string
	Number string
	ExpirationMonth int32
	ExpirationYear int32
	CVV string
	Balance float64
	Limit float64
	CreatedAt time.Time
}

func NewCreditCard() *CreditCard {
	c := &CreditCard{}
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now()
	return c
}