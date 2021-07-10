package dto

import (
	"time"
)

type NewTransaction struct {
	ID string
	OwnerName string
	Number string
	ExpirationMonth int32
	ExpirationYear int32
	CVV string
	Amount float64
	Store string
	Description string
	CreatedAt time.Time
}