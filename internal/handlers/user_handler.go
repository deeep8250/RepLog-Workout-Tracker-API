package handlers

import (
	"fmt"
	"net/http"
	"replog/internal/models"
	"replog/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthServiceInterface
}

func NewAuthHandler(a services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService: a}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing fields",
		})
		return
	}

	if err := h.authService.Register(&req); err != nil {
		fmt.Println("error in service is : ", err.Error())
		if err.Error() == "email exists" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "email exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"response": "Register user successfully",
	})

}

func (h *AuthHandler) LoginUser(c *gin.Context) {

	var User models.LoginRequest
	err := c.ShouldBindJSON(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if User.Email == "" || User.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "empty fields",
		})
		return
	}

	token, err := h.authService.Login(&User)
	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
