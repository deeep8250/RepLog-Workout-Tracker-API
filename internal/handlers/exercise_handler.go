package handlers

import (
	"net/http"
	"replog/internal/models"
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

func (h *UserHandler) CreateExercises(c *gin.Context) {
	var Exercise models.Exercises
	err := c.ShouldBindJSON(&Exercise)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	userIDint := userID.(int)

	Exercise.CreatedBy = &userIDint
	err = h.services.CreateExercises(&Exercise)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "workout successfully created",
	})

}

func (h *UserHandler) CreateWorkouts(c *gin.Context) {

	var workouts models.Workouts
	err := c.ShouldBindJSON(&workouts)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "please provide data for all fields",
		})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	userIDint := userID.(int)
	workouts.UserID = userIDint

	err = h.services.CreateWorkouts(&workouts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cant create the workouts ,something went wrong !",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully cretaed your workouts",
	})

}
