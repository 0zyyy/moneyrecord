package history

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	AddHistory(history History) (History, error)
	UpdateHistory(history History) (History, error)
	FindByIdHistory(ID int) (History, error)
	Delete(ID int) (bool, error)
	FindAll() ([]History, error)
	FindByUserId(ID int) (tx *gorm.DB)
	HistorySearch(ID int, params ...string) ([]ResponseHistory, error)
	IncomeSearch(ID int, params ...string) ([]ResponseHistory, error)
	Detail(ID int, params ...string) (History, error)
	Month() ([]ResponseHistory, error)
	Week() ([]ResponseHistory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// find all
func (r *repository) FindAll() ([]History, error) {
	var histories []History
	err := r.db.Table("history").Find(&histories).Error
	if err != nil {
		return histories, err
	}
	return histories, nil
}

//find by id history

func (r *repository) FindByIdHistory(ID int) (History, error) {
	var history History
	err := r.db.Table("history").Where("id_history = ?", ID).Find(&history).Error
	if err != nil {
		return history, err
	}
	return history, nil
}

// Add history
func (r *repository) AddHistory(history History) (History, error) {
	err := r.db.Table("history").Create(&history).Error
	if err != nil {
		return history, err
	}
	return history, nil
}

// UpdateHistory histroy
func (r *repository) UpdateHistory(history History) (History, error) {
	err := r.db.Table("history").Where("id_history = ?", 1).Save(&history).Error
	if err != nil {
		return history, err
	}
	return history, nil
}

// delete history
func (r *repository) Delete(ID int) (bool, error) {
	db := r.db.Table("history").Where("id_history = ?", ID).Delete(&History{})
	if db.Error != nil {
		return false, db.Error
	} else if db.RowsAffected < 1 {
		return false, errors.New("id not found")
	}
	return true, nil
}

// find by id
func (r *repository) FindByUserId(ID int) (tx *gorm.DB) {
	return r.db.Table("history").Where("id_user = ?", ID)
}

func (r *repository) HistorySearch(ID int, params ...string) ([]ResponseHistory, error) {
	var histories []ResponseHistory
	// params ada 1 otomatis ada date-nya berarti dia ngesearch
	if params[0] != "" {
		log.Println("rintting by date")
		err := r.FindByUserId(ID).Where("date = ?", params[0]).Order("date DESC").Find(&histories).Error
		if err != nil {
			return histories, err
		}
	} else {
		log.Println("Printing all")
		err := r.FindByUserId(ID).Order("date DESC").Find(&histories).Error
		if err != nil {
			return histories, err
		}
	}
	return histories, nil
}

func (r *repository) IncomeSearch(ID int, params ...string) ([]ResponseHistory, error) {
	var histories []ResponseHistory
	fmt.Println(params)
	// params ada 2 otomatis ada date-nya berarti dia ngesearch
	if params[1] != "" {
		err := r.FindByUserId(ID).Where("type = ?", params[0]).Where("date = ?", params[1]).Order("date DESC").Find(&histories).Error
		if err != nil {
			return histories, err
		}
	} else {
		err := r.FindByUserId(ID).Where("type = ?", params[0]).Order("date DESC").Find(&histories).Error
		if err != nil {
			return histories, err
		}
	}
	return histories, nil
}

func (r *repository) Detail(ID int, params ...string) (History, error) {
	var history History
	// find detail history
	err := r.FindByUserId(ID).Where("type = ?", params[0]).Where("date = ?", params[1]).Find(&history).Error
	if err != nil {
		return history, err
	}
	return history, nil
}

func (r *repository) Month(date string) ([]ResponseHistory, error) {
	var histories []ResponseHistory
	err := r.FindByUserId(2).Where("date LIKE ?", "%"+date+"%").Order("date DESC").Find(&histories).Error
	if err != nil {
		return histories, err
	}
	return histories, nil
}

func (r *repository) Week(date string) ([]ResponseHistory, error) {
	var histories []ResponseHistory
	err := r.FindByUserId(2).Where("date >= ?", "%"+date+"%").Order("date DESC").Find(&histories).Error
	if err != nil {
		return histories, err
	}
	return histories, nil
}
