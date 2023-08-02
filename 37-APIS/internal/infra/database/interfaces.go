package database

import "github.com/kthanasio/golang/37-APIS/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(mail string) (*entity.User, error)
}
