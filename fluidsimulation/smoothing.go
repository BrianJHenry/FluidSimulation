package fluidsimulation

import "math"

/*
kernel = (s - d)^3 / integral of that
*/
func DifferenceCubedSmoothingKernel(distance, smoothingRadius float64) float64 {
	if distance >= smoothingRadius {
		return 0
	}

	// Calculate normalization (integral of smoothing function)
	var normalizationFactor = math.Pi * math.Pow(smoothingRadius, 4) / 2

	// Calculate smoothing kernel
	return math.Pow(smoothingRadius-distance, 3) / normalizationFactor
}

func DifferenceCubedSmoothingKernelDerivative(distance, smoothingRadius float64) float64 {
	if distance >= smoothingRadius {
		return 0
	}
	return -6 * math.Pow(smoothingRadius-distance, 2) / (math.Pi * math.Pow(smoothingRadius, 4))
}

/*
kernel = (s^2 - d^2)^3 / integral of that
*/
func SqdDifferenceCubedSmoothingKernel(smoothingRadius, distance float64) float64 {
	if distance >= smoothingRadius {
		return 0
	}

	// Calculate normalization (integral of smoothing function)
	var normalizationFactor = math.Pi * 32 * math.Pow(smoothingRadius, 7) / 35

	// Calculate smoothing kernel
	return math.Pow(math.Pow(smoothingRadius, 2)-math.Pow(distance, 2), 3) / normalizationFactor
}

func SqdDifferenceCubedSmoothingKernelDerivative(smoothingRadius, distance float64) float64 {
	if distance >= smoothingRadius {
		return 0
	}
	return -105 * distance * math.Pow(math.Pow(smoothingRadius, 2)-math.Pow(distance, 2), 2) / (math.Pi * 16 * math.Pow(smoothingRadius, 7))
}

/*
Directly from Sebastion Lague's youtube
*/
func SmoothingKernel(smoothingRadius, distance float64) float64 {
	if distance >= smoothingRadius {
		return 0
	}

	var normalizationFactor = (math.Pi * math.Pow(smoothingRadius, 4)) / 6
	return (smoothingRadius - distance) * (smoothingRadius - distance) / normalizationFactor
}

/*
Directly from Sebastian Lague's youtube
*/
func SmoothingKernelDerivative(smoothingRadius, distance float64) float64 {
	if distance >= smoothingRadius {
		return 0
	}
	return 12 * (distance - smoothingRadius) / (math.Pow(smoothingRadius, 4) * math.Pi)
}
