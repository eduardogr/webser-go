package usecases

import (
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/repositories"
	"github.com/eduardogr/webser-go/internal/adapters/interfaces/usecases"
	"github.com/eduardogr/webser-go/internal/domain"
)

func NewGetAllNumbersUsecase(repository repositories.NumberRepository) usecases.GetAllNumbersUsecase {
	return &GetAllNumbers{
		Repository: repository,
	}
}

type GetAllNumbers struct {
	Repository repositories.GetAllNumbersRepository
}

func (u *GetAllNumbers) Execute() ([]domain.Number, error) {
	return u.Repository.GetAll()
}
