package main

import (
	"math"

	"github.com/woyteck/toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	prevLat  float64
	prevLong float64
}

func NewCalculdatorService() CalculatorServicer {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	var distance float64
	if s.prevLat > 0 && s.prevLong > 0 {
		distance = calculateDistance(data.Lat, data.Long, s.prevLat, s.prevLong)
	}
	s.prevLat = data.Lat
	s.prevLong = data.Long

	return distance, nil
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
