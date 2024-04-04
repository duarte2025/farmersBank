package usecases

import (
	"context"
	"testing"

	"github.com/duarte2025/farmersBank/domain/entities"
)

func TestCreateAccount(t *testing.T) {
	// t.Parallel()

	tests := []struct {
		name  string
		input CreateAccountInput
		setup func(t *testing.T) *CreateAccountUC
		want  *CreateAccountOutput
	}{
		{
			name: "should create account successfully",
			input: CreateAccountInput{
				AccountName: "test",
			},
			setup: func(t *testing.T) *CreateAccountUC {
				return NewCreateAccountUC(&CreateAccountRepositoryMock{
					CreateAccountFunc: func(ctx context.Context, account entities.Account) error {
						return nil
					},
				})
			},
			want: &CreateAccountOutput{
				AccountName: "test",
				AccountID:   "1",
				Balance:     0,
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()

			// PREPARE
			ctx := context.Background()

			// ACT
			got, err := tt.setup(t).Execute(ctx, tt.input.AccountName)

			// ASSERT
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if got.AccountName != tt.want.AccountName {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			if got.Balance != tt.want.Balance {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
