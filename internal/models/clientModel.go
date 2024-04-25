package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id          uuid.UUID `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Address     string    `db:"address"`
	PostalCode  string    `db:"postal_code"`
	City        string    `db:"city"`
	Country     string    `db:"country"`
	Phone       string    `db:"phone"`
	Email       string    `db:"email"`
	DateCreated time.Time `db:"date_created"`
}
