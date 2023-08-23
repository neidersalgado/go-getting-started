package repository_test

import (
	"testing"
	"tu_paquete/models"

	"github.com/DATA-DOG/go-sqlmock"
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
	mock.ExpectExec("INSERT INTO").WithArgs("TestName", "test@example.com", "testPass", "TestAddress", "1990-01-01", "TestCity").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewUserRepo(gormDB)

	user := &models.User{
		Name:      "TestName",
		Email:     "test@example.com",
		Password:  "testPass",
		Address:   "TestAddress",
		Birthdate: "1990-01-01",
		City:      "TestCity",
	}

	err = repo.Create(user)
	assert.NoError(t, err)
}

// Aquí puedes continuar con más tests para las demás funciones...
