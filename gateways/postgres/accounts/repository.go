package accounts

import "github.com/duarte2025/farmersBank/gateways/postgres"

type Repository struct {
	q postgres.Querier
}

func NewRepository(q postgres.Querier) *Repository {
	return &Repository{q: q}
}
