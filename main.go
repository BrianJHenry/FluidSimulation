package main

import (
	"log"

	"github.com/BrianJHenry/FluidSimulation/fluidsimulation"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func (g *Game) Update() error {
	getInputs()
	return fluidsimulation.UpdateParticles()
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, part := range fluidsimulation.Particles {
		vector.DrawFilledCircle(screen, float32(part.Position.X), float32(part.Position.Y), float32(fluidsimulation.ParticleRadius), part.Color, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return fluidsimulation.ScreenWidth, fluidsimulation.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(fluidsimulation.ScreenWidth, fluidsimulation.ScreenHeight)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func init() {
	fluidsimulation.ResetParticles()
}

func getInputs() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		fluidsimulation.Paused = !fluidsimulation.Paused
		println("Pause Toggled")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		fluidsimulation.ResetParticles()
		println("Reset")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		fluidsimulation.IsGravity = !fluidsimulation.IsGravity
		println("Gravity Toggled")
	}
}
