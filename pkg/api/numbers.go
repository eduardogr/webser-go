package api

import (
	"errors"
)

// NumberService contains the methods of the number service
type NumberService interface {
	New(user NewNumberRequest) error
	GetAll() ([]Number, error)
	Get(id int) (Number, error)
	Update(n Number, id int) error
	Remove(id int) error
}

// Number repository is what lets our service do db operations without knowing anything about the implementation
type NumberRepository interface {
	GetAll() ([]Number, error)
	Get(id int) (Number, error)
	Create(n NewNumberRequest) error
	Update(n Number, id int) error
	Remove(id int) error
}

type numberService struct {
	storage NumberRepository
}

func NewNumberService(numberRepo NumberRepository) NumberService {
	return &numberService{
		storage: numberRepo,
	}
}

func (s *numberService) New(n NewNumberRequest) error {
	// validations
	if n.ID < 0 {
		return errors.New("numbers service - negative numbers not allowed.")
	}

	err := s.storage.Create(n)

	if err != nil {
		return err
	}

	return nil
}

func (s *numberService) Get(id int) (Number, error) {
	return s.storage.Get(id)
}

func (s *numberService) GetAll() ([]Number, error) {
	return s.storage.GetAll()
}

func (s *numberService) Update(n Number, id int) error {
	return s.storage.Update(n, id)
}

func (s *numberService) Remove(id int) error {
	return s.storage.Remove(id)
}
