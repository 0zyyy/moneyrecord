package history

type Service interface {
	FindAll() ([]History, error)
}

type service struct {
	histoRepo repository
}

func NewService(histoRepo repository) *service {
	return &service{histoRepo: histoRepo}
}

func (s *service) FindAll() ([]History, error) {
	histories, err := s.histoRepo.FindAll()
	if err != nil {
		return histories, err
	}
	return histories, nil
}
