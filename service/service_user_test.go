package service

import (
	"crud/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserService{Repository: userRepository}

func TestUserService_Get(t *testing.T) {
	userRepository.Mock.On("FindById", "1").Return(nil)

	user, err := userService.Get("1")
	assert.Nil(t, user)
	assert.NotNil(t, err)
}
