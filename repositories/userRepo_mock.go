package repositories

import (
	"backjeep/models"
	"errors"
)

var MockUserDB = make(map[string]*models.User)

type MockUserRepository struct{}

func (m *MockUserRepository) CreateUser(u *models.User) (*models.User, error) {
	if _, exists := MockUserDB[u.Email]; exists {
		return nil, errors.New("user already exists")
	}
	MockUserDB[u.Email] = u
	return u, nil
}

func (m *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	user, exists := MockUserDB[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func init() {
	MockUserDB = make(map[string]*models.User)
}
