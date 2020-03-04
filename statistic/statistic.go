package statistic

import (
	"log"
	"math"
)

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
func Sum(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}
func Quadratic(a, b, c float64) (float64, float64) {
	d := (b * b) - (4 * a * c)
	if d == 0 {
		x1 := -b / (2 * a)
		x2 := x1
		return x1, x2
	}
	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		x1r := math.Round(x1*100) / 100
		x2r := math.Round(x2*100) / 100
		return x1r, x2r
	}
	log.Fatal("D<0")
	return 0, 0
}
