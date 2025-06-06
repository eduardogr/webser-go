package usecases

import "github.com/eduardogr/webser-go/internal/domain"

type CreateNewNumberUsecase interface {
	Execute(n domain.NewNumberRequest) error
}

type GetAllNumbersUsecase interface {
	Execute() ([]domain.Number, error)
}

type GetNumberByIdUsecase interface {
	Execute(id int) (domain.Number, error)
}

type RemoveNumberByIdUsecase interface {
	Execute(id int) error
}

type UpdateNumberByIdUsecase interface {
	Execute(n domain.Number, id int) error
}
