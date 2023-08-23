package repository

import (
	"github.com/heroku/go-getting-started/models"
	"github.com/jinzhu/gorm"
)

type FavoritesRepo struct {
	DB *gorm.DB
}

// NewFavoritesRepo inicializa y devuelve una instancia de FavoritesRepo
func NewFavoritesRepo(DB *gorm.DB) *FavoritesRepo {
	return &FavoritesRepo{DB: DB}
}

// Create inserta un nuevo favorito en la base de datos
func (repo *FavoritesRepo) Create(favorite *models.Favorite) error {
	if err := repo.DB.Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

// Delete elimina un favorito por su ID
func (repo *FavoritesRepo) Delete(id int) error {
	if err := repo.DB.Delete(&models.Favorite{}, id).Error; err != nil {
		return err
	}
	return nil
}
