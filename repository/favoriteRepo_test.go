package repository_test

import (
	"testing"

	"github.com/heroku/go-getting-started/models"
	"github.com/heroku/go-getting-started/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestFavoritesRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error al abrir el mock db %s", err)
	}
	defer db.Close()

	gormDB, _ := gorm.Open("mysql", db)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WithArgs(1, "apiRef123").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := repository.NewFavoritesRepo(gormDB)

	favorite := &models.Favorite{
		UserID: 1,
		RefAPI: "apiRef123",
	}

	err = repo.Create(favorite)
	assert.NoError(t, err)
}
