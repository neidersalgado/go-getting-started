package repository

import (
	"fmt"

	"github.com/heroku/go-getting-started/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

func (repo *UserRepo) Create(user *models.User) error {
	if err := repo.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) Update(user *models.User) error {
	fmt.Printf("user: %v", user)
	if err := repo.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) GetAll() ([]models.User, error) {
	var users []models.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Model(&models.User{}).Select("id,password").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
