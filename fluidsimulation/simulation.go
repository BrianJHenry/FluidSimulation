package fluidsimulation

import (
	"math/rand"
)

const (
	ScreenWidth  = 1080
	ScreenHeight = 720
)

var (
	Particles             = []*Particle{}
	NumberOfParticles     = 500
	BounceDampeningFactor = 0.8
	Paused                = false
	IsGravity             = false
)

func UpdateSimulation() error {
	if !Paused {
		recalculateForces()
		applyAcceleration()
		applyMovement()
	}
	return nil
}

func ResetParticles() {
	Particles = []*Particle{}
	var interval = float64(ScreenWidth) / float64(NumberOfParticles+2)
	for i := 1; i <= NumberOfParticles; i++ {
		var height = rand.Float64()
		var newParticle = Particle{
			Position: Vector2{
				X: float64(i) * interval,
				Y: ScreenHeight * height,
			},
			Speed: Vector2{
				X: 0,
				Y: 0,
			},
		}
		RecalculateForces(0, []*Particle{&newParticle}, ScreenHeight, ScreenWidth, IsGravity)
		Particles = append(Particles, &newParticle)
	}
}

func recalculateForces() {
	for i := range Particles {
		RecalculateForces(i, Particles, ScreenHeight, ScreenWidth, IsGravity)
	}
}

func applyAcceleration() {
	for _, part := range Particles {
		ApplyAcceleration(part)
	}
}

func applyMovement() {
	for _, part := range Particles {
		ApplyMovement(part)
	}
	checkCollideWithEdges()
}

func checkCollideWithEdges() {
	for _, part := range Particles {
		CheckCollisionWithEdges(part, BounceDampeningFactor, ScreenWidth, ScreenHeight)
	}
}
