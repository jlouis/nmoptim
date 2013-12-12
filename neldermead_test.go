package nmoptim

import (
	"testing"
)

func rosen(x []float64) float64 {
	return (100*(x[1]-x[0]*x[0])*(x[1]-x[0]*x[0]) + (1.0-x[0])*(1.0-x[0]))
}

func TestSub(t *testing.T) {
	x := []float64{2.0}
	y := []float64{1.0}
	z := sub(x, y)

	if x[0] != 2.0 {
		t.Errorf("x[0] = %v want %v", x[0], 2.0)
	}

	if y[0] != 1.0 {
		t.Errorf("y[0] = %v want %v", y[0], 1.0)
	}

	if z[0] != 1.0 {
		t.Errorf("z[0] = %v want %v", z[0], 1.0)
	}
}

func TestAdd(t *testing.T) {
	x := []float64{2.0}
	y := []float64{1.0}
	z := add(x, y)

	if x[0] != 2.0 {
		t.Errorf("x[0] = %v want %v", x[0], 2.0)
	}

	if y[0] != 1.0 {
		t.Errorf("y[0] = %v want %v", y[0], 1.0)
	}

	if z[0] != 3.0 {
		t.Errorf("z[0] = %v want %v", z[0], 3.0)
	}
}

func TestScale(t *testing.T) {
	x := point{2.0}

	y := x.scale(2.0)

	if x[0] != 2.0 {
		t.Errorf("x[0] = %v want %v", x[0], 2.0)
	}

	if y[0] != 4.0 {
		t.Errorf("y[0] = %v want %v", y[0], 4.0)
	}
}

func TestCentroid(t *testing.T) {
	t.Errorf("Failing TestCentroid because it is not implemented")
}

func TestArgMax(t *testing.T) {
	t.Errorf("Failing TestArgMax because it is not implemented")
}

func TestArgMin(t *testing.T) {
	t.Errorf("Failing TestArgMin because it is not implemented")
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
