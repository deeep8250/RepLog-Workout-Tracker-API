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

func (h *setHandler) GetAllSets(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorised user",
		})
		return
	}

	workoutID := c.Param("id")
	workoutIDInt, err := strconv.Atoi(workoutID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, err := h.WorkoutService.GetWorkoutsByID(workoutIDInt, userID.(int))
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

	results, err := h.SetService.GetAllSets(workoutIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sets": results,
	})

}

func (h *setHandler) DeleteFromSetsHandler(c *gin.Context) {

	workoutID := c.Param("workoutId")
	workoutIDint, err := strconv.Atoi(workoutID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	setID := c.Param("setId")

	setIDint, err := strconv.Atoi(setID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	UserID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorised users",
		})
		return
	}

	_, err = h.WorkoutService.GetWorkoutsByID(workoutIDint, UserID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.SetService.DeleteFromSets(workoutIDint, setIDint)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "set deleted successfully",
	})

}

func (h *setHandler) ProgressReportHandler(c *gin.Context) {

	userId, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorised user",
		})
		return
	}

	exerciseID := c.Param("exercise_id")
	exerciseIDint, err := strconv.Atoi(exerciseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	value, err := h.SetService.ProgressReportService(userId.(int), exerciseIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"report": value,
	})

}

func (h *setHandler) ShouldIncreaseHandler(c *gin.Context) {

	userId, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorised user",
		})
		return
	}

	targetedReps := c.Query("target_reps")
	target_reps, err := strconv.Atoi(targetedReps)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	value, err := h.SetService.ShouldIncreaseService(target_reps, userId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"report": value,
	})

}
