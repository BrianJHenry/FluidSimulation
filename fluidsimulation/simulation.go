package fluidsimulation

import (
	"image/color"
	"math/rand"
	"sync"

	"github.com/mazznoer/colorgrad"
)

const (
	ScreenWidth  = 1080
	ScreenHeight = 720
)

var (
	Particles                     = []*Particle{}
	NumberOfParticles             = 500
	BounceDampeningFactor         = 0.8
	Paused                        = false
	IsGravity                     = false
	ParticleColors                = colorgrad.Cividis()
	ParticleColorScale    float64 = 7
)

func UpdateSimulation() error {
	if !Paused {
		recalculateForces()
		appyForces()
		applyVelocity()
	}
	return nil
}

func ResetParticles() {
	Particles = []*Particle{}
	var interval = float64(ScreenWidth) / float64(NumberOfParticles+2)
	for i := 0; i < NumberOfParticles; i++ {
		var height = rand.Float64()
		var newParticle = Particle{
			Position: Vector2{
				X: float64(i+1) * interval,
				Y: ScreenHeight * height,
			},
			Velocity: Vector2{
				X: 0,
				Y: 0,
			},
			Color: color.White,
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

func appyForces() {
	var forceWait sync.WaitGroup
	forceWait.Add(len(Particles))
	for _, part := range Particles {
		// Update velocity
		go func(particle *Particle) {
			ApplyForces(particle)
			forceWait.Done()
		}(part)
	}
	forceWait.Wait()
}

func applyVelocity() {
	for _, part := range Particles {
		ApplyVelocity(part)
	}
	checkCollisionWithEdges()
}

func checkCollisionWithEdges() {
	for _, part := range Particles {
		CheckCollisionWithEdges(part, BounceDampeningFactor, ScreenWidth, ScreenHeight)
	}
}
