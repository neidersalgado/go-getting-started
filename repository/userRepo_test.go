package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/heroku/go-getting-started/models"
	"github.com/heroku/go-getting-started/repository"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error al abrir el mock db %s", err)
	}
	defer db.Close()

	gormDB, _ := gorm.Open("mysql", db)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WithArgs("TestName", "test@example.com", "testPass", "TestAddress", sqlmock.AnyArg(), "TestCity").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := repository.NewUserRepo(gormDB)
	dateTime, _ := time.Parse("2006-01-02", "1990-01-01")
	user := &models.User{
		Name:      "TestName",
		Email:     "test@example.com",
		Password:  "testPass",
		Address:   "TestAddress",
		Birthdate: dateTime,
		City:      "TestCity",
	}

	err = repo.Create(user)
	assert.NoError(t, err)
}
