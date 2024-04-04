package entities

import "github.com/google/uuid"

type Account struct {
	ID      uuid.UUID
	Name    string
	Balance int
}

func NewAccount(id uuid.UUID, name string, balance int) Account {
	return Account{
		ID:      id,
		Name:    name,
		Balance: balance,
	}
}
