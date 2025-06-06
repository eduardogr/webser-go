package usecases

import (
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

func NewUpdateNumberByIdUsecase(repository repositories.NumberRepository) usecases.UpdateNumberByIdUsecase {
	return &UpdateNumberById{
		Repository: repository,
	}
}

type UpdateNumberById struct {
	Repository repositories.UpdateNumberById
}

func (u *UpdateNumberById) Execute(n domain.Number, id int) error {
	return u.Repository.Update(n, id)
}
