package datastructures_test

import (
	"MarluxGitHub/adventOfCode/pkg/datastructures"
	"testing"
)

func TestMooreDistance(t *testing.T) {
	tests := []struct {
		name     string
		point1   datastructures.Point
		point2   datastructures.Point
		expected int
	}{
		{
			name:     "Test 1: Points with positive coordinates",
			point1:   datastructures.Point{X: 1, Y: 1},
			point2:   datastructures.Point{X: 4, Y: 5},
			expected: 4,
		},
		{
			name:     "Test 2: Points with negative coordinates",
			point1:   datastructures.Point{X: -1, Y: -1},
			point2:   datastructures.Point{X: -4, Y: -5},
			expected: 4,
		},
		// {
		// 	name:     "Test 3: Points with mixed positive and negative coordinates",
		// 	point1:   datastructures.Point{X: -1, Y: 1},
		// 	point2:   datastructures.Point{X: 4, Y: -5},
		// 	expected: 9,
		// },
		{
			name:     "Test 4: Identical points",
			point1:   datastructures.Point{X: 1, Y: 1},
			point2:   datastructures.Point{X: 1, Y: 1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.point1.MooreDistance(tt.point2); got != tt.expected {
				t.Errorf("MooreDistance() = %v, want %v", got, tt.expected)
			}
		})
	}
}
