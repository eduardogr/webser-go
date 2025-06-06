package usecases

import (
	"errors"

	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

func NewCreateNumberUsecase(repository repositories.NumberRepository) usecases.CreateNewNumberUsecase {
	return &CreateNewNumber{
		Repository: repository,
	}
}

type CreateNewNumber struct {
	Repository repositories.CreateNumberRepository
}

func (u *CreateNewNumber) Execute(n domain.NewNumberRequest) error {
	// validations
	if n.ID < 0 {
		return errors.New("numbers service - negative numbers not allowed")
	}

	err := u.Repository.Create(n)

	if err != nil {
		return err
	}

	return nil
}
