package fluidsimulation

import (
	"image/color"
	"math/rand"
	"sync"

	"github.com/mazznoer/colorgrad"
)

const (
	ScreenWidth   = 1080
	ScreenHeight  = 720
	ParticleCount = 200
)

var (
	Particles                           = [ParticleCount]*Particle{}
	Densities                           = [ParticleCount]float64{}
	BounceDampeningFactor               = 0.8
	Paused                              = false
	HasGravity                          = false
	ParticleColors                      = colorgrad.Cividis()
	ParticleRadius                      = 7
	ParticleRadiusOfInfluence   float64 = 100
	BoundaryDistanceOfInfluence float64 = 25
	ParticleColorScale          float64 = 7
)

func UpdateParticles() error {
	if !Paused {
		if HasGravity {
			applyGravity()
		}
		calculateDensities()
		calculateVelocityGradientsForParticles()
		applyColorToParticles()
		applyVelocityToParticles()
	}
	return nil
}

func ResetParticles() {
	var interval = float64(ScreenWidth) / float64(ParticleCount+2)
	for i := 0; i < ParticleCount; i++ {
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
		Particles[i] = &newParticle
	}
}

func applyGravity() {
	for _, part := range Particles {
		part.Velocity = part.Velocity.Add(Vector2{
			X: 0,
			Y: 0.2,
		})
	}
}

func calculateDensities() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(ParticleCount)
	for i, part := range Particles {
		go func(i int, part *Particle) {
			Densities[i] = CalculateDensity(part)
			waitGroup.Done()
		}(i, part)
	}
	waitGroup.Wait()
}

func calculateVelocityGradientsForParticles() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(ParticleCount)
	for i, part := range Particles {
		go func(i int, part *Particle) {
			var velocityGradient = CalculatePressureGradient(i)
			part.Velocity = part.Velocity.Add(velocityGradient)
			waitGroup.Done()
		}(i, part)
	}
	waitGroup.Wait()
}

func applyColorToParticles() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(ParticleCount)
	for _, part := range Particles {
		func(part *Particle) {
			applyColor(part)
			waitGroup.Done()
		}(part)
	}
	waitGroup.Wait()
}

func applyVelocityToParticles() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(ParticleCount)
	for _, part := range Particles {
		func(part *Particle) {
			applyVelocity(part)
			checkCollisionWithEdges(part, BounceDampeningFactor, ScreenWidth, ScreenHeight)
			waitGroup.Done()
		}(part)
	}
	waitGroup.Wait()
}
