package usecase

import (
	entity "crud/entity/requests"
	"crud/entity/responses"
	"crud/repository"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type Usecase interface {
	CreateUser(user *entity.User) (*entity.User, error)
	Login(user entity.Login) (*responses.UserLogin, error)
	FindUser(id string) (*entity.User, error)
	FindUserx(id string) (*entity.Userx, error)
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
	NewUser.Password = string(HashPassword)
	User, err := u.repository.Create(NewUser)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (u *usecase) FindUser(id string) (*entity.User, error) {
	ids, errs := strconv.Atoi(id)
	if errs != nil {
		return nil, errs
	}
	User, err := u.repository.GetId(ids)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (u *usecase) FindUserx(id string) (*entity.Userx, error) {
	ids, errs := strconv.Atoi(id)
	if errs != nil {
		return nil, errs
	}
	User, err := u.repository.GetIdx(ids)
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
	HashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(HashPassword)
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
	resp.Name = User.Name
	resp.Age = User.Age
	resp.Nasabah = User.Nasabah
	resp.Email = User.Email
	return &resp, nil
}

func (u *usecase) Login(user entity.Login) (*responses.UserLogin, error) {
	input := user
	input.Email = user.Email
	input.Password = user.Password

	newuser, err := u.repository.GetByEmail(input.Email)
	check := bcrypt.CompareHashAndPassword([]byte(newuser.Password), []byte(input.Password))
	if err != check {
		resp := responses.UserLogin{}
		return &resp, err
	} else {
		resp := responses.UserLogin{}
		resp.ID = newuser.ID
		resp.CreatedAt = newuser.CreatedAt
		resp.UpdatedAt = newuser.UpdatedAt
		resp.Name = newuser.Name
		resp.Age = newuser.Age
		resp.Nasabah = newuser.Nasabah
		resp.Email = newuser.Email
		//resp.Token = newuser.Token
		return &resp, nil
	}
}
