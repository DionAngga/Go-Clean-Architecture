package repository

import (
	"crud/entity"
)

type UserRepository interface {
	FindById(id string) *entity.User
}
