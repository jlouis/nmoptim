package nmoptim

import (
	"math"
	"testing"
)

const (
	precision = 0.00001
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
	sx := simplex{point{1.0, 0.0}, point{0.0, 1.0}, point{3.0, 4.0}}
	p := sx.centroid(2)

	if sx[0][0] != 1.0 && sx[0][1] != 0.0 {
		t.Errorf("sx[0] is %v", sx[0])
	}

	if sx[1][0] != 0.0 && sx[0][1] != 1.0 {
		t.Errorf("sx[1] is %v", sx[1])
	}

	if sx[2][0] != 3.0 && sx[2][1] != 4.0 {
		t.Errorf("sx[2] is %v", sx[2])
	}

	if p[0] != 0.5 && p[1] != 0.5 {
		t.Errorf("p is %v want %v", p, point{0.5, 0.5})
	}
}

func TestNelderMead(t *testing.T) {
	x := []float64{-0.5, -1.0}
	y := []float64{-0.75, 1.5}
	z := []float64{0.2, 1.2}

	s := [][]float64{x, y, z}

	r := Optimize(rosen, s)
	if math.Abs(r[0]-1.0) > precision {
		t.Errorf("r[0] = %v, want %v", r[0], 1.0)
	}

	if math.Abs(r[1]-1.0) > precision {
		t.Errorf("r[1] = %v, want %v", r[1], 1.0)
	}
}

func BenchmarkNelderMead(b *testing.B) {
		x := []float64{-0.5, -1.0}
	y := []float64{-0.75, 1.5}
	z := []float64{0.2, 1.2}

	s := [][]float64{x, y, z}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Optimize(rosen, s)
	}
}