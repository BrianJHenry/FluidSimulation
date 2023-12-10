package fluidsimulation

const threshold = 0.000001

func AlmostEqual(value1, value2 float64) bool {
	var difference = value1 - value2
	if difference < 0 {
		difference *= -1
	}
	return difference <= threshold
}
