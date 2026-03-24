package handlers

import (
	"net/http"
	"replog/internal/models"
	"replog/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WorkoutHandler struct {
	services *services.WorkoutsServices
}

func NewWorkoutHandler(Services *services.WorkoutsServices) *WorkoutHandler {
	return &WorkoutHandler{
		services: Services,
	}
}

func (h *WorkoutHandler) CreateWorkouts(c *gin.Context) {

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

func (h *WorkoutHandler) GetAllWorkouts(c *gin.Context) {

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	userIDint := userID.(int)
	workouts, err := h.services.GetWorkouts(userIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"workouts": workouts,
	})

}

func (h *WorkoutHandler) GetWorkoutByID(c *gin.Context) {

	Id := c.Param("id")
	if Id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "please provide the value",
		})
		return
	}

	IDint, err := strconv.Atoi(Id)
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

	workout, err := h.services.GetWorkoutsByID(IDint, userIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid input",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"workout": workout,
	})

}

func (h *WorkoutHandler) DeleteWorkoutByID(c *gin.Context) {

	Id := c.Param("id")
	if Id == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "please provide the workout id",
		})
		return
	}

	IDint, err := strconv.Atoi(Id)
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

	err = h.services.DeleteWorkouts(IDint, userIDint)
	if err != nil {

		if err.Error() == "workout not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "workout not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "workout deleted",
	})

}
