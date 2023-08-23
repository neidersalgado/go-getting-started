package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/heroku/go-getting-started/controllers/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/models"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockUserRepo)
	mockToken := new(mocks.MockTokenGenerator)
	controller := NewUserController(mockRepo, mockToken)

	r := gin.Default()
	r.POST("/register", controller.Register)

	user := models.User{Email: "test@email.com", Password: "testPassword"}
	mockRepo.On("Create", &models.User{Email: "test@email.com", Password: "testPassword"}).Return(nil)
	mockToken.On("GenerateToken", mock.AnythingOfType("string")).Return("someToken", nil)

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	mockRepo.AssertExpectations(t)
	mockToken.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockUserRepo)
	mockToken := new(mocks.MockTokenGenerator)
	controller := NewUserController(mockRepo, mockToken)

	user := models.User{Email: "login@email.com", Password: "loginPass"}
	mockRepo.On("GetUserByEmail", user.Email).Return(&user, nil)
	mockToken.On("GenerateToken", mock.AnythingOfType("string")).Return("someToken", nil)

	r := gin.Default()
	r.POST("/login", controller.Login)

	loginInfo := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "login@email.com",
		Password: "loginPass",
	}

	body, _ := json.Marshal(loginInfo)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	mockRepo.AssertExpectations(t)
	mockToken.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := mocks.MockUserRepo{}
	mockTokenGen := &mocks.MockTokenGenerator{}
	controller := NewUserController(&mockRepo, mockTokenGen)

	// Assume this user is registered and trying to update.
	user := models.User{Email: "update@email.com", Password: "updatePass"}
	mockRepo.On("Create", &user).Return(nil)
	mockRepo.On("Update", mock.Anything).Return(nil)

	mockToken := "mocked.jwt.token"
	mockTokenGen.On("GenerateToken", mock.Anything).Return(mockToken, nil)
	mockTokenGen.On("ValidateToken", mockToken).Return(user.Email, nil)

	r := gin.Default()
	r.PUT("/update", controller.UpdateUser)

	updateInfo := models.User{
		Email: "update@email.com",
		Name:  "New Name",
	}

	body, _ := json.Marshal(updateInfo)
	req, _ := http.NewRequest(http.MethodPut, "/update", bytes.NewBuffer(body))

	// Add the mocked JWT to the request header
	req.Header.Add("Authorization", mockToken)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
