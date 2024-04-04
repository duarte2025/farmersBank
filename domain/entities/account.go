package entities

import "github.com/google/uuid"

type Account struct {
	id      uuid.UUID
	name    string
	balance int
}

func NewAccount(name string) Account {
	return Account{
		id:      uuid.New(),
		name:    name,
		balance: 0,
	}
}

func (a Account) ID() string {
	return a.id.String()
}

func (a Account) Name() string {
	return a.name
}

func (a Account) Balance() int {
	return a.balance
}
