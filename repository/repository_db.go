package repository

import (
	entity "crud/entity/requests"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	Get(user *[]entity.User) (*[]entity.User, error)
	GetId(id int) (*entity.User, error)
	GetIdx(id int) (*entity.Userx, error)
	Create(user *entity.User) (*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(user *entity.User, id string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

const DNS = "root:@tcp(localhost:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

func (db *repository) Get(user *[]entity.User) (*[]entity.User, error) {
	err := db.DB.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *repository) GetId(id int) (*entity.User, error) {
	var user = &entity.User{}
	err := db.DB.Where("id = ?", id).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *repository) GetIdx(id int) (*entity.Userx, error) {
	var user = &entity.Userx{}
	err := db.DB.Where("id = ?", id).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *repository) Create(user *entity.User) (*entity.User, error) {
	err := db.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *repository) Update(user *entity.User) (*entity.User, error) {
	err := db.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *repository) Delete(user *entity.User, id string) (*entity.User, error) {
	err := db.DB.Delete(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *repository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := db.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}
