package fluidsimulation

import "math"

type Particle struct {
	Position, Speed Vector2
	Forces          []Vector2
}

func (part *Particle) ApplyAcceleration() {
	for _, force := range part.Forces {
		part.Speed = part.Speed.Add(force)
	}
}

func (part *Particle) ApplyMovement() {
	part.Position = part.Position.Add(part.Speed)
}

func (part *Particle) CheckCollisionWithEdges(bounceDampeningFactor, screenWidth, screenHeight float64) {
	if part.Position.X > screenWidth {
		part.Position.X = screenWidth
		part.Speed.X *= -1 * bounceDampeningFactor
	} else if part.Position.X < 0 {
		part.Position.X = 0
		part.Speed.X *= -1 * bounceDampeningFactor
	}
	if part.Position.Y > screenHeight {
		part.Position.Y = screenHeight
		part.Speed.Y *= -1 * bounceDampeningFactor
	} else if part.Position.Y < 0 {
		part.Position.Y = 0
		part.Speed.Y *= -1 * bounceDampeningFactor
	}
}

func (part *Particle) RecalculateForces(parts []*Particle, height, width float64, gravity bool) {
	part.Forces = []Vector2{}
	if gravity {
		part.Forces = append(part.Forces, Vector2{
			X: 0,
			Y: 0.2,
		})
	}
	part.CalculateInterParticleRepulsion(parts)
	part.getBoundaryRepulsion()
}

func (part *Particle) CalculateInterParticleRepulsion(parts []*Particle) {
	for _, p := range parts {
		var distance = part.Position.GetDistance(p.Position)
		if distance > 0.000001 {
			var scalingFactor = getForceScalingFactorFromDistance(distance)
			var forceVector = part.Position.GetUnitDirection(p.Position).ScalarMultiply(scalingFactor)
			part.Forces = append(part.Forces, forceVector)
		}
	}
}

func getForceScalingFactorFromDistance(distance float64) float64 {
	var closeness = 100 - distance
	if closeness < 0 {
		return 0
	} else {
		return math.Pow(closeness, 3) / 1000000
	}
}

func (part *Particle) getBoundaryRepulsion() {
	if part.Position.X < 25 {
		part.Forces = append(part.Forces, Vector2{
			X: getForceScalingFactorFromDistance(part.Position.X * 4),
			Y: 0,
		})
	} else if 1080-part.Position.X < 25 {
		part.Forces = append(part.Forces, Vector2{
			X: -getForceScalingFactorFromDistance((1080 - part.Position.X) * 4),
			Y: 0,
		})
	}
	if part.Position.Y < 25 {
		part.Forces = append(part.Forces, Vector2{
			X: 0,
			Y: getForceScalingFactorFromDistance(part.Position.Y * 4),
		})
	} else if 720-part.Position.Y < 25 {
		part.Forces = append(part.Forces, Vector2{
			X: 0,
			Y: -getForceScalingFactorFromDistance((720 - part.Position.Y) * 4),
		})
	}
}
