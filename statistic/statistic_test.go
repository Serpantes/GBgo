package statistic

import "testing"

type testpair struct {
	values []float64
	result float64
}
type testQuad struct {
	a  float64
	b  float64
	c  float64
	x1 float64
	x2 float64
}

var testAvg = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}
var testSum = []testpair{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

var testQuads = []testQuad{
	{3.2, -7.8, 1, 2.30, 0.14},
	{2, 4, 2, -1, -1},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range testAvg {
		v := Average(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
func TestSum(t *testing.T) {
	for _, pair := range testSum {
		v := Sum(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
func TestQ(t *testing.T) {
	for _, nums := range testQuads {
		v1, v2 := Quadratic(nums.a, nums.b, nums.c)
		if v1 != nums.x1 || v2 != nums.x2 {
			t.Error(
				"For", nums.a, nums.b, nums.c,
				"expected", nums.x1, nums.x2,
				"got", v1, v2,
			)
		}
	}
}
