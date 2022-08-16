package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	AddUser(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindAll() ([]User, error)
	FindById(userId int) (User, error)
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

func (r *repository) AddUser(user User) (User, error) {
	err := r.db.Table("user").Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Table("user").Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindById(userId int) (User, error) {
	var user User
	err := r.db.Table("user").Where("id_user = ?", userId).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
