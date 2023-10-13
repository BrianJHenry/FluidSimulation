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
			Y: 0.5,
		})
	}
	part.CalculateInterParticleRepulsion(parts)
	part.CalculateBoundaryRepulsion(height, width)
}

func (part *Particle) CalculateInterParticleRepulsion(parts []*Particle) {
	for _, p := range parts {
		var distance = part.Position.GetDistance(p.Position)
		if distance > 0.000001 {
			var forceVector = part.Position.GetUnitDirection(p.Position).ScalarDivide(math.Pow(part.Position.GetDistance(p.Position), 3))
			part.Forces = append(part.Forces, forceVector)
		}
	}
}

func (part *Particle) CalculateBoundaryRepulsion(height, width float64) {
	var verticalForceFactor = height/2 - part.Position.Y
	var horizontalForceFactor = width/2 - part.Position.X
	part.Forces = append(part.Forces, Vector2{
		X: 1 / horizontalForceFactor,
		Y: 1 / verticalForceFactor,
	})
}
