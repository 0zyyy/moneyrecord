package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var user []User
	err := r.db.Table("user").Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
