package usecase

import (
	"crud/entity"
	"crud/repository"
)

func Get(user *[]entity.User) {
	repository.DB.Find(user)
}

func GetId(users *entity.User, id string) {
	repository.DB.First(users, id)
}

func Create(user *entity.User) {
	repository.DB.Create(user)
}

func Update(user *entity.User) {
	repository.DB.Save(&user)
}

func Deleted(user *entity.User, id string) {
	repository.DB.Delete(user, id)
}
