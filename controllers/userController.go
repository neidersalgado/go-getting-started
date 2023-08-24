package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/models"
)

// Define the user repository interface
type UserRepo interface {
	Create(user *models.User) error
	Update(user *models.User) error
	GetAll() ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

// Define the token generator interface
type TokenGenerator interface {
	GenerateToken(email string) (string, error)
	ValidateToken(tokenstring string) (string, error)
}

// UserController structure that houses the user repository and token generator
type UserController struct {
	Repo     UserRepo
	TokenGen TokenGenerator
}

// NewUserController returns a new UserController instance
func NewUserController(r UserRepo, tg TokenGenerator) *UserController {
	return &UserController{Repo: r, TokenGen: tg}
}

// Register handles user registration
func (uc *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.Repo.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating the user", "msg": err.Error()})
		return
	}

	token, err := uc.TokenGen.GenerateToken(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating the token", "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}

// UpdateUser handles updating a user's information
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := ctx.GetHeader("Authorization") // Asumiendo que el token se envía en el header de Autorización
	emailFromToken, err := uc.TokenGen.ValidateToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "msg": err})
		return
	}

	if emailFromToken != user.Email {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission to update this user", "msg": emailFromToken, "userEmail": user.Email})
		return
	}

	err = uc.Repo.Update(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the user", "msg": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "User updated successfully"})
}

// Login handles user authentication
func (uc *UserController) Login(ctx *gin.Context) {
	var loginInfo struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pwd, err := uc.Repo.GetUserByEmail(loginInfo.Email)
	if err != nil || pwd != loginInfo.Email {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password", "msg": loginInfo, "err": err.Error()})
		return
	}

	token, err := uc.TokenGen.GenerateToken(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating the token", "msg": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
