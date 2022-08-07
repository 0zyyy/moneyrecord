package user

type Service interface {
	FindAll() ([]User, error)
}

type service struct {
	UserRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.UserRepository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}
