package fluidsimulation

import "math"

type Vector2 struct {
	X, Y float64
}

func (v1 *Vector2) Add(v2 Vector2) Vector2 {
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

func (v1 Vector2) ScalarDivide(scalar float64) Vector2 {
	return Vector2{
		X: v1.X / scalar,
		Y: v1.Y / scalar,
	}
}

func (v1 Vector2) ScalarMultiply(scalar float64) Vector2 {
	return Vector2{
		X: v1.X * scalar,
		Y: v1.Y * scalar,
	}
}

func (v1 *Vector2) GetUnitDirection(v2 Vector2) Vector2 {
	var unnormalized = v1.Subtract(v2)
	var unnormalizedSqd = unnormalized.Squared()
	var mag = math.Sqrt(unnormalizedSqd.X + unnormalizedSqd.Y)
	return unnormalized.ScalarDivide(mag)
}

func (v1 *Vector2) GetDistance(v2 Vector2) float64 {
	var differenceVector = v1.Subtract(v2)
	var differenceVectorSquared = differenceVector.Squared()
	return math.Sqrt(differenceVectorSquared.X + differenceVectorSquared.Y)
}
