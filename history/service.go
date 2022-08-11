package history

type Service interface {
	AddHistory(history History) (History, error)
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

func (s *service) AddHistory(history History) (History, error) {
	newHistory, err := s.histoRepo.AddHistory(history)
	if err != nil {
		return newHistory, err
	}
	return newHistory, nil
}
