package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMax(t *testing.T) {

	testCases := []struct {
		Name   string
		Min    float64
		Max    float64
		Values []float64
		Result []float64
	}{
		{
			Name:   "¿Qué debería pasar si pasa correctamente un min, un máximo y un conjunto de valores?",
			Min:    6,
			Max:    12,
			Values: []float64{52, 8, 10, 1, 30, 54},
			Result: []float64{8, 10},
		},
		{
			Name:   "¿Qué debería pasar si solo pasas un min, un máximo y un solo valor?",
			Min:    3,
			Max:    65,
			Values: []float64{9},
			Result: []float64{9},
		},
		{
			Name:   "¿Qué debería suceder si el valor pasado como min es realmente mayor que max?",
			Min:    8,
			Max:    2,
			Values: []float64{1, 4, 9},
			Result: nil,
		},
		{
			Name:   "¿Qué debería suceder si no se pasa un solo valor dentro del rango?",
			Min:    3,
			Max:    15,
			Values: []float64{},
			Result: nil,
		},
		{
			Name:   "¿Qué debería pasar si min y max son negativos?",
			Min:    -2,
			Max:    -10,
			Values: []float64{1, 68, 58, 5},
			Result: nil,
		},
		{
			Name:   "¿Qué debería pasar si min y max son iguales?",
			Min:    7,
			Max:    7,
			Values: []float64{8, 2, 87, 6},
			Result: nil,
		},
	}

	for _, testCase := range testCases {
		result := minMax(testCase.Min, testCase.Max, testCase.Values...)
		assert.Equal(t, testCase.Result, result)
	}
}
