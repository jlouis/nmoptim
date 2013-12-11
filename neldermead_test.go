package nmoptim

import (
	"testing"
)

func rosen(x []float64) float64 {
	return (100*(x[1]-x[0]*x[0])*(x[1]-x[0]*x[0]) + (1.0-x[0])*(1.0-x[0]))
}

func TestNelderMead(t *testing.T) {
	x := []float64{-0.5, -1.0}
	y := []float64{-0.75, 1.5}
	z := []float64{0.2, 1.2}

	s := [][]float64{x, y, z}

	r := Optimize(rosen, s)
	if r[0] != 1.0 {
		t.Errorf("r[0] = %v, want %v", r[0], 1.0)
	}

	if r[1] != 1.0 {
		t.Errorf("r[1] = %v, want %v", r[1], 1.0)
	}
}
