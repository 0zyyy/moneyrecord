package history

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	AddHistory(history History) (History, error)
	Update(history History) (History, error)
	Delete(ID int) (bool, error)
	FindAll() ([]History, error)
	FindById(ID int) ([]History, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]History, error) {
	var histories []History
	err := r.db.Table("history").Find(&histories).Error
	if err != nil {
		return histories, err
	}
	return histories, nil
}

// Add history
func (r *repository) AddHistory(history History) (History, error) {
	err := r.db.Table("history").Create(&history).Error
	if err != nil {
		return history, err
	}
	return history, nil
}

// update histroy
func (r *repository) Update(history History) (History, error) {
	err := r.db.Save(&history).Error
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

// find by id tapi bentukane array
func (r *repository) FindById(ID int) ([]History, error) {
	var histories []History
	err := r.db.Table("history").Where("id_user = ?", ID).Order("date DESC").Find(&histories).Error
	if err != nil {
		return histories, err
	}
	return histories, nil
}
