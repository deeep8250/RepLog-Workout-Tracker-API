package handlers

import (
	"net/http"
	"replog/internal/models"
	"replog/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(AuthService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: AuthService}
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
		if err.Error() == "email already exist" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "email already exists",
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
