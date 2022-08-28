package history

import (
	"log"
	"strconv"
	"time"

	"github.com/0zyyy/money_record/helper"
	"github.com/hyperjiang/php"
)

type Service interface {
	Create(input NewHistoryInput) (History, error)
	FindAll() ([]History, error)
	SearchHistory(ID int, date string) ([]ResponseHistory, error)
	SearchIncome(ID int, tipe string, date string) ([]ResponseHistory, error)
	Analysis(idUser int, date string) (ResponseAnalysis, error)
	Update(input NewHistoryInput) (History, error)
	Delete(IDHistory int) (bool, error)
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

func (s *service) Create(input NewHistoryInput) (History, error) {
	history := History{
		IDUser:    input.IDUser,
		Type:      input.Type,
		Total:     input.Total,
		Date:      input.Date,
		Details:   input.Details,
		CreatedAt: time.Now().Format("2006-01-02"),
		UpdatedAt: time.Now().Format("2006-01-02"),
	}
	newHistory, err := s.histoRepo.AddHistory(history)
	if err != nil {
		return newHistory, err
	}
	return newHistory, nil
}

func (s *service) SearchHistory(ID int, date string) ([]ResponseHistory, error) {
	// find by id dulu
	log.Println(date)
	history, err := s.histoRepo.HistorySearch(ID, date)
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

func (s *service) Analysis(idUser int, date string) (ResponseAnalysis, error) {
	var monthIncome, monthOutcome float64
	today, err := php.DateCreate(date)
	if err != nil {
		return ResponseAnalysis{}, err
	}
	thisMonth := php.DateFormat(today, "Y-m")
	resultMonth, err := s.histoRepo.Month(idUser, thisMonth)
	if err != nil {
		return ResponseAnalysis{}, err
	}
	week := []string{} // buat date week
	weekly := []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	// thisMonth := php.DateFormat(today, "Y-m")
	// bikin day

	// bikin week susah bhangsat
	week = append(week, php.DateFormat(today, "Y-m-d"))
	for i := 1; i <= 6; i++ {
		dateInterval, err := php.DateIntervalCreateFromDateString("1 day") // make date interval 1 day sebelumnya
		if err != nil {
			return ResponseAnalysis{}, err
		}
		daySebelum, err := php.DateCreate(week[i-1]) // init daySebelum
		if err != nil {
			return ResponseAnalysis{}, err
		}
		week = append(week, php.DateFormat(php.DateSub(daySebelum, dateInterval), "Y-m-d")) // append date sebelumnya, 7 hari kebelakang
	}
	week = helper.Reverse(week)
	resultWeek, err := s.histoRepo.Week(idUser, week[0])
	if err != nil {
		return ResponseAnalysis{}, err
	}
	for i := 0; i <= len(resultWeek)-1; i++ {
		if resultWeek[i].Type == "Pengeluaran" {
			for j := 0; j < len(week); j++ {
				if resultWeek[i].Date == week[j] {
					convertedTotal, err := strconv.ParseFloat(resultWeek[i].Total, 64) // convert string of total to float64
					if err != nil {
						panic(err)
					}
					weekly[j] = convertedTotal
				}
			}
		}
	}
	for i := 0; i <= len(resultMonth)-1; i++ {
		if resultMonth[i].Type == "Pemasukan" {
			convertedInc, err := strconv.ParseFloat(resultMonth[i].Total, 64)
			if err != nil {
				panic(err)
			}
			monthIncome += convertedInc
		} else {
			convertedOut, err := strconv.ParseFloat(resultMonth[i].Total, 64)
			if err != nil {
				panic(err)
			}
			monthOutcome += convertedOut
		}
	}
	return ResponseAnalysis{
		Today:     weekly[6],
		Yesterday: weekly[5],
		Week:      weekly,
		Month: MonthResult{
			Income:  monthIncome,
			Outcome: monthOutcome,
		},
	}, nil
}

func (s *service) Update(input NewHistoryInput) (History, error) {
	history, err := s.histoRepo.FindByIdHistory(input.IDHistory)
	if err != nil {
		return history, err
	}
	history.IDUser = input.IDUser
	history.Date = input.Date
	history.Details = input.Details
	history.Total = input.Total
	history.Type = input.Type
	history.UpdatedAt = time.Now().Format("2006-01-02")

	updatedHistory, err := s.histoRepo.UpdateHistory(history)
	if err != nil {
		return updatedHistory, err
	}
	return updatedHistory, nil
}

func (s *service) Delete(IDHistory int) (bool, error) {
	deleted, err := s.histoRepo.Delete(IDHistory)
	if err != nil {
		return deleted, err
	}
	return deleted, nil
}
