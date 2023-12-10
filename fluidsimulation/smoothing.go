package fluidsimulation

import "math"

/*
kernel = (s - d)^3
*/
func differenceCubedSmoothingKernel(distance, smoothingRadius float64) float64 {
	var closeness = smoothingRadius - distance
	if closeness <= 0 {
		return 0
	}

	// Calculate normalization (integral of smoothing function)
	var normalizationFactor = math.Pi * math.Pow(smoothingRadius, 4) / 2

	// Calculate smoothing kernel
	return math.Pow(closeness, 3) / normalizationFactor
}

/*
kernel = (s^2 - d^2)^3
*/
func sqdDifferenceCubedSmoothingKernel(smoothingRadius, distance float64) float64 {
	if smoothingRadius-distance <= 0 {
		return 0
	}
	var closeness = (math.Pow(smoothingRadius, 2) - math.Pow(distance, 2))

	// Calculate normalization (integral of smoothing function)
	var normalizationFactor = math.Pi * 32 * math.Pow(smoothingRadius, 7) / 35

	// Calculate smoothing kernel
	return math.Pow(closeness, 3) / normalizationFactor
}
