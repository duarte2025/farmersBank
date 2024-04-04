package usecases

import (
	"context"
	"fmt"

	"github.com/duarte2025/farmersBank/domain/entities"
)

type CreateAccountRepository interface {
	CreateAccount(ctx context.Context, account entities.Account) error
}

type CreateAccountUC struct {
	repository CreateAccountRepository
}

func NewCreateAccountUC(repository CreateAccountRepository) *CreateAccountUC {
	return &CreateAccountUC{
		repository: repository,
	}
}

type CreateAccountInput struct {
	AccountName string
}

type CreateAccountOutput struct {
	AccountName string
	AccountID   string
	Balance     int
}

func (uc *CreateAccountUC) Execute(ctx context.Context, accountName string) (*CreateAccountOutput, error) {
	// TODO create telemetry in project
	// const spanName = "CreateAccountUC.Execute"

	// ctx, span := telemetryfs.FromContext(ctx).Start(ctx, spanName)
	// defer span.End()

	account := entities.NewAccount(accountName)
	err := uc.repository.CreateAccount(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %w", err)
	}

	return &CreateAccountOutput{
		AccountName: account.Name(),
		AccountID:   account.ID(),
		Balance:     account.Balance(),
	}, nil
}
