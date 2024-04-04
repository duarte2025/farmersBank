package entities

import "github.com/google/uuid"

type Entry struct {
	ID        uuid.UUID
	Operation string
	AccountID uuid.UUID
	Amount    int
	Version   int
}

func NewEntry(id, accountID uuid.UUID, operation string, amount int) Entry {
	return Entry{
		ID:        id,
		Operation: operation,
		AccountID: accountID,
		Amount:    amount,
	}
}
