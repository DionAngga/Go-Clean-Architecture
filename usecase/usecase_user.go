package usecase

import (
	"crud/entity"
	"crud/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Usecase interface {
	CreateUser(user *entity.User) *entity.User
	//Login(user *entity.User) (*entity.User, error)
	FindUser(user *entity.User, id string) (*entity.User, error)
	FindAllUser(user *[]entity.User) (*[]entity.User, error)
	FindUserByEmail(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User, id string) (*entity.User, error)
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(repository repository.Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) CreateUser(user *entity.User) *entity.User {
	NewUser := user
	NewUser.Name = user.Name
	NewUser.Age = user.Age
	NewUser.Nasabah = user.Nasabah
	NewUser.Email = user.Email
	HashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	fmt.Println(string(HashPassword))
	NewUser.Password = string(HashPassword)

	u.repository.Create(NewUser)

	return NewUser
}

// func (u *usecase) Login(user *entity.User) (*entity.User, error) {
// 	input := user
// 	input.Email = user.Email
// 	input.Password = user.Password

// 	user, err := u.repository.GetByEmail(input.Email)
// 	if err != nil {
// 		return user, err
// 	}
// 	if user.ID == 0 {
// 		return user, errors.New("user not found")
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

func (u *usecase) FindUser(user *entity.User, id string) (*entity.User, error) {
	User, err := u.repository.GetId(user, id)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (u *usecase) FindAllUser(user *[]entity.User) (*[]entity.User, error) {
	User, err := u.repository.Get(user)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (u *usecase) UpdateUser(user *entity.User) (*entity.User, error) {
	User, err := u.repository.Update(user)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (u *usecase) DeleteUser(user *entity.User, id string) (*entity.User, error) {
	User, err := u.repository.Delete(user, id)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (u *usecase) FindUserByEmail(user *entity.User) (*entity.User, error) {
	email := user.Email
	User, err := u.repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return User, nil
}
