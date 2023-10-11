package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

type particle struct {
	xPos, yPos     float64
	xSpeed, ySpeed float64
}

const (
	screenWidth, screenHeight = 640, 480
)

var (
	firstParticle particle
)

func (g *Game) Update() error {
	getInputs()
	applyMovement()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(firstParticle.xPos), float32(firstParticle.yPos), 16, color.White, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func init() {
	firstParticle = particle{
		xPos:   screenWidth / 2,
		yPos:   screenHeight / 2,
		xSpeed: 0,
		ySpeed: 0,
	}
}

func getInputs() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		firstParticle.ySpeed += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		firstParticle.ySpeed -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		firstParticle.xSpeed += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		firstParticle.xSpeed -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		firstParticle.xSpeed = 0
		firstParticle.ySpeed = 0
	}
}

func applyMovement() {
	firstParticle.xPos += firstParticle.xSpeed
	firstParticle.yPos += firstParticle.ySpeed
	checkCollideWithEdges()
}

func checkCollideWithEdges() {
	if firstParticle.xPos > screenWidth || firstParticle.xPos < 0 {
		firstParticle.xSpeed *= -1
		firstParticle.xPos += firstParticle.xSpeed
	}
	if firstParticle.yPos > screenHeight || firstParticle.yPos < 0 {
		firstParticle.ySpeed *= -1
		firstParticle.yPos += firstParticle.ySpeed
	}
}
