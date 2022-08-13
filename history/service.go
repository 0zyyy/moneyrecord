package history

type Service interface {
	AddHistory(history History) (History, error)
	FindAll() ([]History, error)
	SearchHistory(ID int, tipe string, date string) ([]ResponseHistory, error)
	SearchIncome(ID int, tipe string, date string) ([]ResponseHistory, error)
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

func (s *service) SearchHistory(ID int, tipe string, date string) ([]ResponseHistory, error) {
	// find by id dulu
	history, err := s.histoRepo.HistorySearch(ID, tipe, date)
	if err != nil {
		return history, err
	}
	return history, nil
}

func (s *service) SearchIncome(ID int, tipe string, date string) ([]ResponseHistory, error) {
	// find by id dulu
	history, err := s.histoRepo.IncomeSearch(ID, tipe, date)
	if err != nil {
		return history, err
	}
	return history, nil
} 