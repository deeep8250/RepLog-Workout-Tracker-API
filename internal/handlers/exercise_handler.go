package handlers

import (
	"net/http"
	"replog/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	services *services.ExerciseServices
}

func NewExerciseHandler(Services *services.ExerciseServices) *UserHandler {
	return &UserHandler{
		services: Services,
	}
}

func (h *UserHandler) GetAllExercises(c *gin.Context) {

	UserId, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user isnt found",
		})
		return
	}

	muscle := c.Query("muscle")

	exercises, err := h.services.GetAllExercises(muscle, UserId.(int))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exercises": exercises,
	})

}
