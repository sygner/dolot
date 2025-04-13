package utils

import "math"

func ComputeRank(givenImpression uint32) uint32 {
	currentFactor := math.Log(float64(givenImpression)) / math.Log(1.5)
	roundedFactor := math.Round(currentFactor)
	if roundedFactor == 0 {
		roundedFactor = 1
	}
	return uint32(roundedFactor)
}
