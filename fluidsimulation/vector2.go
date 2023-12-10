package fluidsimulation

import (
	"errors"
	"math"
)

type Vector2 struct {
	X, Y float64
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 Vector2) Subtract(v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func (v1 Vector2) Squared() Vector2 {
	return Vector2{
		X: v1.X * v1.X,
		Y: v1.Y * v1.Y,
	}
}

func (v1 Vector2) ScalarDivide(scalar float64) (Vector2, error) {
	if AlmostEqual(scalar, 0) {
		return Vector2{
			X: 0,
			Y: 0,
		}, errors.New("No division by 0.")
	}
	return Vector2{
		X: v1.X / scalar,
		Y: v1.Y / scalar,
	}, nil
}

func (v1 Vector2) ScalarMultiply(scalar float64) Vector2 {
	return Vector2{
		X: v1.X * scalar,
		Y: v1.Y * scalar,
	}
}

func (v Vector2) Magnitude() float64 {
	var sqdV = v.Squared()
	return math.Sqrt(sqdV.X + sqdV.Y)
}

func (v1 Vector2) GetUnitDirection(v2 Vector2) Vector2 {
	var unnormalized = v1.Subtract(v2)
	var mag = unnormalized.Magnitude()
	unitDirection, _ := unnormalized.ScalarDivide(mag)
	return unitDirection
}

func (v1 Vector2) GetDistance(v2 Vector2) float64 {
	var differenceVector = v1.Subtract(v2)
	return differenceVector.Magnitude()
}
