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
	Particles                           = []*Particle{}
	NumberOfParticles                   = 500
	BounceDampeningFactor               = 0.8
	Paused                              = false
	IsGravity                           = false
	ParticleColors                      = colorgrad.Cividis()
	ParticleRadius                      = 7
	ParticleRadiusOfInfluence   float64 = 100
	BoundaryDistanceOfInfluence float64 = 25
	ParticleColorScale          float64 = 7
)

func UpdateParticles() error {
	if !Paused {
		recalculateForcesForParticles()
		applyForcesToParticles()
		applyColorToParticles()
		applyVelocityToParticles()
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

func recalculateForcesForParticles() {
	for i := range Particles {
		RecalculateForces(i, Particles, ScreenHeight, ScreenWidth, IsGravity)
	}
}

func applyForcesToParticles() {
	var forceWait sync.WaitGroup
	forceWait.Add(len(Particles))
	for _, part := range Particles {
		// Update velocity
		go func(particle *Particle) {
			applyForces(particle)
			forceWait.Done()
		}(part)
	}
	forceWait.Wait()
}

func applyColorToParticles() {
	for _, part := range Particles {
		applyColor(part)
	}
}

func applyVelocityToParticles() {
	for _, part := range Particles {
		applyVelocity(part)
	}
	checkCollisionWithEdgesForParticles()
}

func checkCollisionWithEdgesForParticles() {
	for _, part := range Particles {
		checkCollisionWithEdges(part, BounceDampeningFactor, ScreenWidth, ScreenHeight)
	}
}
