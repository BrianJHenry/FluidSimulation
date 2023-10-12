package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

type particle struct {
	xPos, yPos     float64
	xSpeed, ySpeed float64
	forces         []vector2
}

type vector2 struct {
	x, y float64
}

const (
	screenWidth  = 1080
	screenHeight = 720
)

var (
	particles             = []*particle{}
	numberOfParticles     = 40
	bounceDampeningFactor = 0.8
	paused                = false
)

func (g *Game) Update() error {
	getInputs()
	if !paused {
		applyAcceleration()
		applyMovement()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, part := range particles {
		vector.DrawFilledCircle(screen, float32(part.xPos), float32(part.yPos), 8, color.White, false)
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
	particles = []*particle{}
	var interval = float64(screenWidth) / float64(numberOfParticles+2)
	var height = screenHeight * 0.25
	for i := 1; i <= numberOfParticles+1; i++ {
		particles = append(particles, &particle{
			xPos:   float64(i) * interval,
			yPos:   height,
			xSpeed: 0,
			ySpeed: 0,
			forces: []vector2{
				{
					x: 0,
					y: 0.5,
				},
			},
		})
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
}

func applyAcceleration() {
	for _, part := range particles {
		for _, force := range part.forces {
			part.xSpeed += force.x
			part.ySpeed += force.y
		}
	}
}

func applyMovement() {
	for _, part := range particles {
		part.xPos += part.xSpeed
		part.yPos += part.ySpeed
	}
	checkCollideWithEdges()
}

func checkCollideWithEdges() {
	for _, part := range particles {
		if part.xPos > screenWidth {
			part.xPos = screenWidth
			part.xSpeed *= -1 * bounceDampeningFactor
		} else if part.xPos < 0 {
			part.xPos = 0
			part.xSpeed *= -1 * bounceDampeningFactor
		}
		if part.yPos > screenHeight {
			part.yPos = screenHeight
			part.ySpeed *= -1 * bounceDampeningFactor
		} else if part.yPos < 0 {
			part.yPos = 0
			part.ySpeed *= -1 * bounceDampeningFactor
		}
	}
}
