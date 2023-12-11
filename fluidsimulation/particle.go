package fluidsimulation

import (
	"image/color"
)

type Particle struct {
	Position, Velocity Vector2
	Color              color.Color
}

func applyVelocity(part *Particle) {
	part.Position = part.Position.Add(part.Velocity)
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

func checkCollisionWithEdges(part *Particle, bounceDampingFactor, screenWidth, screenHeight float64) {
	if part.Position.X > screenWidth {
		part.Position.X = screenWidth
		part.Velocity.X *= -1 * bounceDampingFactor
	} else if part.Position.X < 0 {
		part.Position.X = 0
		part.Velocity.X *= -1 * bounceDampingFactor
	}
	if part.Position.Y > screenHeight {
		part.Position.Y = screenHeight
		part.Velocity.Y *= -1 * bounceDampingFactor
	} else if part.Position.Y < 0 {
		part.Position.Y = 0
		part.Velocity.Y *= -1 * bounceDampingFactor
	}
}

// func getBoundaryRepulsion(part *Particle) {
// 	if part.Position.X < BoundaryDistanceOfInfluence {
// 		part.Forces = append(part.Forces, Vector2{
// 			X: sqdDifferenceCubedSmoothingKernel(BoundaryDistanceOfInfluence, part.Position.X),
// 			Y: 0,
// 		})
// 	} else if ScreenWidth-part.Position.X < BoundaryDistanceOfInfluence {
// 		part.Forces = append(part.Forces, Vector2{
// 			X: -sqdDifferenceCubedSmoothingKernel(BoundaryDistanceOfInfluence, ScreenWidth-part.Position.X),
// 			Y: 0,
// 		})
// 	}
// 	if part.Position.Y < BoundaryDistanceOfInfluence {
// 		part.Forces = append(part.Forces, Vector2{
// 			X: 0,
// 			Y: sqdDifferenceCubedSmoothingKernel(BoundaryDistanceOfInfluence, part.Position.Y),
// 		})
// 	} else if ScreenHeight-part.Position.Y < BoundaryDistanceOfInfluence {
// 		part.Forces = append(part.Forces, Vector2{
// 			X: 0,
// 			Y: -sqdDifferenceCubedSmoothingKernel(BoundaryDistanceOfInfluence, ScreenHeight-part.Position.Y),
// 		})
// 	}
// }
