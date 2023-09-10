package repositories

import (
	"backjeep/models"
	"backjeep/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.CreateUserRequest) (models.User, error)
}

type UserRepo struct{}

func (ur *UserRepo) GetAllUsers() ([]models.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := utils.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepo) CreateUser(req models.CreateUserRequest) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email"
	var user models.User

	err = utils.DB.QueryRow(query, req.Name, req.Email, hashedPassword).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
