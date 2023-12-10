package fluidsimulation

import (
	"image/color"
	"math"
)

type Particle struct {
	Position, Velocity Vector2
	Forces             []Vector2
	Color              color.Color
}

func ApplyForces(part *Particle) {
	for _, force := range part.Forces {
		part.Velocity = part.Velocity.Add(force)
	}
	applyColor(part)
}

func ApplyVelocity(part *Particle) {
	part.Position = part.Position.Add(part.Velocity)
}

func CheckCollisionWithEdges(part *Particle, bounceDampeningFactor, screenWidth, screenHeight float64) {
	if part.Position.X > screenWidth {
		part.Position.X = screenWidth
		part.Velocity.X *= -1 * bounceDampeningFactor
	} else if part.Position.X < 0 {
		part.Position.X = 0
		part.Velocity.X *= -1 * bounceDampeningFactor
	}
	if part.Position.Y > screenHeight {
		part.Position.Y = screenHeight
		part.Velocity.Y *= -1 * bounceDampeningFactor
	} else if part.Position.Y < 0 {
		part.Position.Y = 0
		part.Velocity.Y *= -1 * bounceDampeningFactor
	}
}

func RecalculateForces(index int, particles []*Particle, height, width float64, gravity bool) {
	var part = particles[index]
	part.Forces = []Vector2{}
	if gravity {
		part.Forces = append(part.Forces, Vector2{
			X: 0,
			Y: 0.2,
		})
	}
	CalculateInterParticleRepulsion(index, particles)
	getBoundaryRepulsion(part)
}

func CalculateInterParticleRepulsion(index int, particles []*Particle) {
	var part = particles[index]
	for i, p := range particles {
		if i != index {
			var distance = part.Position.GetDistance(p.Position)
			var scalingFactor = getForceScalingFactorFromDistance(distance)
			var forceVector = part.Position.GetUnitDirection(p.Position).ScalarMultiply(scalingFactor)
			part.Forces = append(part.Forces, forceVector)
		}
	}
}

func applyColor(part *Particle) {
	var scaledVel = part.Velocity.Magnitude() / ParticleColorScale
	if scaledVel > 1.0 {

		scaledVel = 1.0
	} else if scaledVel < 0.0 {
		scaledVel = 0.0
	}
	part.Color = ParticleColors.At(scaledVel)
}

func getForceScalingFactorFromDistance(distance float64) float64 {
	var closeness = 100 - distance
	if closeness < 0 {
		return 0
	} else {
		return math.Pow(closeness, 3) / 1000000
	}
}

func getBoundaryRepulsion(part *Particle) {
	if part.Position.X < 25 {
		part.Forces = append(part.Forces, Vector2{
			X: getForceScalingFactorFromDistance(part.Position.X * 4),
			Y: 0,
		})
	} else if ScreenWidth-part.Position.X < 25 {
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
	} else if ScreenHeight-part.Position.Y < 25 {
		part.Forces = append(part.Forces, Vector2{
			X: 0,
			Y: -getForceScalingFactorFromDistance((720 - part.Position.Y) * 4),
		})
	}
}
