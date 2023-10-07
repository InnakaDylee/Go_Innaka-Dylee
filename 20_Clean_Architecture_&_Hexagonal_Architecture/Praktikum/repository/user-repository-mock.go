package repository

import (
	"Praktikum/entity"
	"errors"

	mock "github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	Mock mock.Mock
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

var data = []entity.User{
	{
		ID:       1,
		Name:     "Innaka",
		Email:    "Innaka@gmail.com",
		Password: "Password123",
	},
	{
		ID:       2,
		Name:     "Dylee",
		Email:    "Dylee@gmail.com",
		Password: "Password321"},
}

func (m *MockUserRepository) CreateUserController(user *entity.User) error {

	argument := m.Mock.Called(user)

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("All fields must be filled")
	}

	if argument.Get(0) == nil {
		return errors.New("error")
	} else {
		return nil
	}
}

func (m *MockUserRepository) GetUsersController() (error, interface{}) {
	if data == nil {
		return errors.New("empty data"), nil
	} else {
		return nil, data
	}
}

func (*MockUserRepository) LoginUserController(user *entity.User) (error, string) {

	if user.Email == "" && user.Password == "" {
		return errors.New("All fields must be filled"), ""
	}

	var index int
	found := 0
	for i, val := range data {
		if val.Email == user.Email {
			index = i
			found = 1
			break
		}
	}

	if found == 0 {
		return errors.New("user not found"), ""
	}

	if data[index].Password != user.Password {
		return errors.New("login failed"), ""
	}

	*user = data[index]

	return nil, ""
}