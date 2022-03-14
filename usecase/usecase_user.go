package usecase

import (
	entity "crud/entity/requests"
	"crud/entity/responses"
	"crud/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Usecase interface {
	CreateUser(user *entity.User) (*entity.User, error)
	Login(user entity.Login) (*responses.UserRespon, error)
	FindUser(user *entity.User, id string) (*entity.User, error)
	FindAllUser(user *[]entity.User) (*[]entity.User, error)
	FindUserByEmail(user *entity.Login) (*responses.UserRespon, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User, id string) (*entity.User, error)
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(repository repository.Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) CreateUser(user *entity.User) (*entity.User, error) {
	NewUser := user
	NewUser.Name = user.Name
	NewUser.Age = user.Age
	NewUser.Nasabah = user.Nasabah
	NewUser.Email = user.Email
	HashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	fmt.Println(string(HashPassword))
	NewUser.Password = string(HashPassword)
	User, err := u.repository.Create(NewUser)
	if err != nil {
		return nil, err
	}
	return User, nil
}

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

func (u *usecase) FindUserByEmail(user *entity.Login) (*responses.UserRespon, error) {
	email := user.Email
	User, err := u.repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	resp := responses.UserRespon{}
	resp.ID = User.ID
	resp.CreatedAt = User.CreatedAt
	resp.UpdatedAt = User.UpdatedAt
	resp.DeletedAt = User.DeletedAt
	resp.Name = User.Name
	resp.Age = User.Age
	resp.Nasabah = User.Nasabah
	resp.Email = User.Email
	return &resp, nil
}

func (u *usecase) Login(user entity.Login) (*responses.UserRespon, error) {
	input := user
	input.Email = user.Email
	input.Password = user.Password

	newuser, err := u.repository.GetByEmail(input.Email)
	check := bcrypt.CompareHashAndPassword([]byte(newuser.Password), []byte(input.Password))
	if err != check {
		resp := responses.UserRespon{}
		return &resp, err
	} else {
		resp := responses.UserRespon{}
		resp.ID = newuser.ID
		resp.CreatedAt = newuser.CreatedAt
		resp.UpdatedAt = newuser.UpdatedAt
		resp.DeletedAt = newuser.DeletedAt
		resp.Name = newuser.Name
		resp.Age = newuser.Age
		resp.Nasabah = newuser.Nasabah
		resp.Email = newuser.Email
		return &resp, nil
	}
}
