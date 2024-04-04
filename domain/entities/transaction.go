package entities

import "github.com/google/uuid"

type Transaction struct {
	ID      uuid.UUID
	Entries []Entry
}
