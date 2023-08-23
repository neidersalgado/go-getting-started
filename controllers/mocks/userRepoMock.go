package mocks

import (
	"github.com/heroku/go-getting-started/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) Update(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) GetAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepo) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

type MockTokenGenerator struct {
	mock.Mock
}

func (m *MockTokenGenerator) GenerateToken(email string) (string, error) {
	args := m.Called(email)
	return args.String(0), args.Error(1)
}

func (m *MockTokenGenerator) ValidateToken(tokenstring string) (string, error) {
	args := m.Called(tokenstring)
	return args.String(0), args.Error(1)
}
