package space

import "math"

type Planet string

var ratios = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const SecondsInYear float64 = 31557600.0

func Age(sec float64, planet Planet) float64 {
	cal := sec * 100 / SecondsInYear / ratios[planet]
	cal = float64(int(math.Round(cal))) / 100.00
	return cal
}
