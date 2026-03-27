package services

import (
	"fmt"
	"replog/internal/models"
	"strconv"
)

func Calculate1RM(weight, reps float64) int {

	rm := weight * (1 + reps/30)
	return int(rm)

}

func CalculateTotalVolume(sets []models.Sets) int {

	var totalVolumes float64
	for _, value := range sets {
		WeightInt, err := strconv.ParseFloat(value.Weight_kg, 64)
		if err != nil {
			fmt.Println(err.Error())
			continue

		}

		tempV := WeightInt * float64(value.Reps)
		totalVolumes += float64(tempV)

	}

	return int(totalVolumes)
}

func ShouldIncreaseWeight(sets []models.Sets, targetReps int) bool {

	if len(sets) < 3 {
		return false
	}
	for _, value := range sets {

		if value.Reps < targetReps {
			return false
		}

	}
	return true

}
