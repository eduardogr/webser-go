package repositories

import "github.com/eduardogr/webser-go/internal/domain"

type NumberRepository interface {
	Initialize() error
	CloseConnections()

	GetAllNumbersRepository
	GetNumberById
	CreateNumberRepository
	UpdateNumberById
	RemoveNumberById
}

type GetAllNumbersRepository interface {
	GetAll() ([]domain.Number, error)
}

type GetNumberById interface {
	Get(id int) (domain.Number, error)
}

type CreateNumberRepository interface {
	Create(n domain.NewNumberRequest) error
}

type UpdateNumberById interface {
	Update(n domain.Number, id int) error
}

type RemoveNumberById interface {
	Remove(id int) error
}
