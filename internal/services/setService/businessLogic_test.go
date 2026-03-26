package services

import (
	"replog/internal/models"
	"testing"
)

func TestCalculate1Rm(t *testing.T) {
	tests := []struct {
		name     string
		weight   float64
		reps     float64
		expected int
	}{
		{name: "norma case", weight: 80, reps: 10, expected: 106},
		{name: "zero reps", weight: 80, reps: 0, expected: 80},
		{name: "zero weight", weight: 0, reps: 10, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := Calculate1RM(tt.weight, tt.reps)
			if value != tt.expected {
				t.Errorf("expected %d got %d", tt.expected, value)
			}
		})
	}
}

func TestTotalVolume(t *testing.T) {
	tests := []struct {
		Name     string
		Sets     []models.Sets
		expected int
	}{
		{Name: "normal case", Sets: []models.Sets{
			{Reps: 10, Weight_kg: "80"},
			{Reps: 8, Weight_kg: "80"},
		},
			expected: 1440,
		},
		{
			Name:     "empty sets",
			Sets:     []models.Sets{},
			expected: 0,
		},
		{
			Name: "non numeric weight",
			Sets: []models.Sets{
				{Reps: 10, Weight_kg: "body weight"},
			},
			expected: 0, // skipped
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			value := CalculateTotalVolume(tt.Sets)
			if value != tt.expected {
				t.Errorf("expected %d got %d", tt.expected, value)
			}

		})

	}
}

func TestShouldIncreaseWeight(t *testing.T) {
	tests := []struct {
		Name         string
		Sets         []models.Sets
		TargetedReps int
		expected     bool
	}{

		{
			Name: "no need to increse", Sets: []models.Sets{
				{Reps: 8},
				{Reps: 10},
				{Reps: 10},
			},
			TargetedReps: 10,
			expected:     false,
		},
		{
			Name: " need to increse", Sets: []models.Sets{
				{Reps: 10},
				{Reps: 10},
				{Reps: 10},
			},
			TargetedReps: 10,
			expected:     true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {
			value := ShouldIncreaseWeight(tt.Sets, tt.TargetedReps)
			if value != tt.expected {
				t.Errorf("expected %t got %t", tt.expected, value)
			}

		})

	}
}
