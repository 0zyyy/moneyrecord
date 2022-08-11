package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input LoginInput) (User, error)
	RegisterUser(input RegisterUserInput) (User, error)
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

func (s *service) Login(input LoginInput) (User, error) {
	user, err := s.UserRepository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}
	if user.IDUser == 0 {
		return user, errors.New("no user found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, errors.New("username atau password salah")
	}
	return user, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Nama
	user.Email = input.Email
	user.CreatedAt = time.Now().Local().Format("2006-01-02 15:04:05")
	user.UpdatedAt = time.Now().Local().Format("2006-01-02 15:04:05")
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	newUser, err := s.UserRepository.AddUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
