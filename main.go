package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/BrianJHenry/FluidSimulation/fluidsimulation"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

const (
	screenWidth  = 1080
	screenHeight = 720
)

var (
	particles             = []*fluidsimulation.Particle{}
	numberOfParticles     = 20
	bounceDampeningFactor = 0.8
	paused                = false
	isGravity             = false
)

func (g *Game) Update() error {
	getInputs()
	if !paused {
		recalculateForces()
		applyAcceleration()
		applyMovement()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, part := range particles {
		vector.DrawFilledCircle(screen, float32(part.Position.X), float32(part.Position.Y), 8, color.White, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func resetParticles() {
	particles = []*fluidsimulation.Particle{}
	var interval = float64(screenWidth) / float64(numberOfParticles+2)
	for i := 1; i <= numberOfParticles; i++ {
		var height = rand.Float64()
		var newParticle = fluidsimulation.Particle{
			Position: fluidsimulation.Vector2{
				X: float64(i) * interval,
				Y: screenHeight * height,
			},
			Speed: fluidsimulation.Vector2{
				X: 0,
				Y: 0,
			},
		}
		newParticle.RecalculateForces([]*fluidsimulation.Particle{}, screenHeight, screenWidth, isGravity)
		particles = append(particles, &newParticle)
	}
}

func init() {
	resetParticles()
}

func getInputs() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		paused = !paused
		println("Pause Toggled")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		resetParticles()
		println("Reset")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		isGravity = !isGravity
		println("Gravity Toggled")
	}
}

func recalculateForces() {
	for _, part := range particles {
		part.RecalculateForces(particles, screenHeight, screenWidth, isGravity)
	}
}

func applyAcceleration() {
	for _, part := range particles {
		part.ApplyAcceleration()
	}
}

func applyMovement() {
	for _, part := range particles {
		part.ApplyMovement()
	}
	checkCollideWithEdges()
}

func checkCollideWithEdges() {
	for _, part := range particles {
		part.CheckCollisionWithEdges(bounceDampeningFactor, screenWidth, screenHeight)
	}
}
