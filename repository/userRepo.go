package repository

import (
	"tu_paquete/models" // Reemplaza "tu_paquete" con el nombre de tu paquete

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

// NewUserRepo inicializa y devuelve una instancia de UserRepo
func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

// Create crea un nuevo usuario en la base de datos
func (repo *UserRepo) Create(user *models.User) error {
	if err := repo.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Update actualiza la información del usuario en la base de datos
func (repo *UserRepo) Update(user *models.User) error {
	if err := repo.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// GetAll retorna una lista de todos los usuarios
func (repo *UserRepo) GetAll() ([]models.User, error) {
	var users []models.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
