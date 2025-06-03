package main

import (
	"testing"
)

func TestMinMax(t *testing.T) {
	cases := []struct {
		name     string
		min, max float64
		values   []float64
		expected []float64
	}{
		{
			name:     "1. what should happen if you correctly pass a min, a max and a set of values?",
			min:      2,
			max:      65,
			values:   []float64{7, 69, 45, 25, 1},
			expected: []float64{7, 45, 25},
		},
		{
			name:     "2. What should happen if the you just pass a min, a max and a single value?",
			min:      2,
			max:      65,
			values:   []float64{7},
			expected: []float64{7},
		},
		{
			name:     "3. What should happen if the value passed as min is actually greater than max?",
			min:      8,
			max:      2,
			values:   []float64{3, 6, 12},
			expected: []float64{},
		},
		{
			name:     "4. What should happen if not a single value passed is within the range?",
			min:      2,
			max:      65,
			values:   []float64{1, 99, 999},
			expected: []float64{},
		},
		{
			name:     "5. What should happen if both min and max are negative?",
			min:      -8,
			max:      -150,
			values:   []float64{8, 9, 87},
			expected: []float64{},
		},
		{
			name:     "6. What should happen if both min and max are equal?",
			min:      9,
			max:      9,
			values:   []float64{2, 658, 25},
			expected: []float64{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := minMax(c.min, c.max, c.values...)

			if len(result) != len(c.expected) {
				t.Errorf("fail [%s]: Expected %v, but it was obtained %v, range: [%.f, %.f]", c.name, c.expected, result, c.min, c.max)
				return
			}

			for i := range result {
				if result[i] != c.expected[i] {
					t.Errorf("fail [%s]: expected %v in position %d of slice, but obtained %v. Filter values: [%.f y %.f]", c.name, c.expected[i], i, result[i], c.min, c.max)
				}
			}
		})
	}

}
