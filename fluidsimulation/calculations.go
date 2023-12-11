package fluidsimulation

import "math/rand"

var (
	TargetDensity      float64 = 3
	PressureMultiplier float64 = 5
)

func CalculateDensity(particle *Particle) float64 {
	var density float64 = 0
	var mass float64 = 1
	for _, p := range Particles {
		var distance = particle.Position.GetDistance(p.Position)
		var scalingFactor = SmoothingKernel(ParticleRadiusOfInfluence, distance)
		density += scalingFactor * mass
	}
	return density
}

func CalculatePressureGradient(index int) Vector2 {
	var pressureGradient = Vector2{
		X: 0,
		Y: 0,
	}

	var particle = Particles[index]

	for i, p := range Particles {
		if i == index {
			continue
		}
		var distance = particle.Position.GetDistance(p.Position)
		var scalingFactor = SmoothingKernelDerivative(ParticleRadiusOfInfluence, distance)
		var gradientDirection = particle.Position.GetUnitDirection(p.Position)
		if AlmostEqual(gradientDirection.X, 0) && AlmostEqual(gradientDirection.Y, 0) {
			gradientDirection = getRandomDirection()
		}
		var sharedPressure = CalculateSharedPressure(Densities[index], Densities[i])
		var gradient = gradientDirection.ScalarMultiply(sharedPressure * scalingFactor / Densities[i])
		pressureGradient = pressureGradient.Add(gradient)
	}

	return pressureGradient
}

func getRandomDirection() Vector2 {
	var vector = Vector2{
		X: rand.Float64() - 0.5,
		Y: rand.Float64() - 0.5,
	}
	return vector.GetUnitDirection(Vector2{
		X: 0,
		Y: 0,
	})
}

func CalculateSharedPressure(density1, density2 float64) float64 {
	return CalculatePressureFromDensity(density1) + CalculatePressureFromDensity(density2)/2
}

func CalculatePressureFromDensity(density float64) float64 {
	var error = density - TargetDensity
	var pressure = error * PressureMultiplier
	return pressure
}
