package history

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]History, error)
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
