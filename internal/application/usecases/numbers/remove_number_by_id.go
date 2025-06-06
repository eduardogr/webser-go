package usecases

import (
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
)

func NewRemoveNumberByIdUsecase(repository repositories.NumberRepository) usecases.RemoveNumberByIdUsecase {
	return &RemoveNumberById{
		Repository: repository,
	}
}

type RemoveNumberById struct {
	Repository repositories.RemoveNumberById
}

func (u *RemoveNumberById) Execute(id int) error {
	return u.Repository.Remove(id)
}
