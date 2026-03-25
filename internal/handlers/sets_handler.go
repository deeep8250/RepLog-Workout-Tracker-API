package handlers

import (
	"net/http"
	"replog/internal/models"
	"replog/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type setHandler struct {
	SetService     services.SetsServiceInterface
	WorkoutService services.WorkoutServiceInterface
}

func NewSetHandler(set services.SetsServiceInterface, workout services.WorkoutServiceInterface) *setHandler {
	return &setHandler{
		SetService: set,

		WorkoutService: workout,
	}
}

func (h *setHandler) CreateSets(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthroized user",
		})
		return

	}

	workoutID := c.Param("id")
	workoutIDint, err := strconv.Atoi(workoutID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid parameter",
		})
		return
	}

	var UserInput models.Sets
	err = c.ShouldBindJSON(&UserInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	t, err := h.WorkoutService.GetWorkoutsByID(workoutIDint, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "workouts not belongs to the user",
		})
		return
	}

	UserInput.WorkoutID = workoutIDint

	err = h.SetService.CreateSets(&UserInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "set created",
	})

}
