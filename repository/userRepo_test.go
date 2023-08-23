package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/heroku/go-getting-started/models"
	"github.com/heroku/go-getting-started/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepo_Create(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error open mock db %s", err)
	}
	defer db.Close()

	// GORM setup with custom dialector
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("Error setting up gorm DB: %s", err)
	}

	// SQLMock expectations
	mock.ExpectBegin()
	mock.ExpectQuery(`^INSERT INTO "users" \("name","email","password","address","birthdate","city"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6\) RETURNING "id"$`).
		WithArgs("TestName", "test@example.com", "testPass", "TestAddress", sqlmock.AnyArg(), "TestCity").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
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
