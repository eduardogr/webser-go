package usecases

import (
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

func NewGetNumberByIdUsecase(repository repositories.NumberRepository) usecases.GetNumberByIdUsecase {
	return &GetNumberById{
		Repository: repository,
	}
}

type GetNumberById struct {
	Repository repositories.GetNumberById
}

func (u *GetNumberById) Execute(id int) (domain.Number, error) {
	return u.Repository.Get(id)
}
