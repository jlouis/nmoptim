package nmoptim

import (
	"fmt"
	"math"
)

const (
	kMax = 1000    // arbitrarily chosen value for now
	ε    = 0.00001 // Stopping criterion point
	α    = 1.0
	β    = 0.5
	γ    = 2.0
)

// point is the type of points in ℝ^n
type point []float64

// simplex is the type used to represent a simplex
type simplex []point

// optfunc is the type of optimization functions. They run from ℝ^n → ℝ, here represented with float64's
type optfunc func([]float64) float64

// Optimize function f with Nelder-Mead. start points to a slice of starting points
// It is the responsibility of the caller to make sure the dimensionality is correct.
func Optimize(f optfunc, start [][]float64) []float64 {
	n := len(start)
	c := len(start[0])
	points := make([]point, 0)
	for _, p := range start {
		points = append(points, point(p))
	}
	sx := simplex(points)
	if n != c+1 {
		panic("Can't optimize with too few starting points")
	}

	var l int
	k := 0
	for ; stopCriterion(f, sx) && k < kMax; k++ {
		h := sx.argMax(f)
		l = sx.argMin(f)

		xp := sub(
			sx.centroid(h).scale(1.0+α),
			sx[h].scale(α))

		if f(xp) < f(sx[l]) {
			xpp := sub(xp.scale(1.0-γ), sx.centroid(h).scale(γ))
			if f(xpp) < f(sx[l]) {
				fmt.Printf("Expanding⋯\n")
				sx[h] = xpp // Expansion
			} else {
				fmt.Printf("Reflecting⋯\n")
				sx[h] = xp // Reflection
			}
		} else if testForallBut(f, xp, sx, h) {
			if f(xp) <= f(sx[h]) {
				sx[h] = xp
			}

			xpp := add(sx[h].scale(β), sx.centroid(h).scale(1.0-β))
			if f(xpp) > f(sx[h]) {
				// Multiple contraction
				fmt.Printf("Multiple contracting⋯\n")
				for i := range sx {
					sx[i] = add(sx[i], sx[l]).scale(0.5)
				}
			} else {
				fmt.Printf("Contracting⋯\n")
				sx[h] = xpp // Contraction
			}
		} else {
			fmt.Printf("Reflecting (2)⋯\n")
			sx[h] = xp // Reflection
		}
	}

	fmt.Printf("Exited after %v iterations\n", k)
	return sx[l]
}

// sub perform point subtraction
func sub(x point, y point) point {
	r := x
	for i := range y {
		r[i] -= y[i]
	}

	return r
}

// add perform point addition
func add(x point, y point) point {
	r := x
	for i := range y {
		r[i] += y[i]
	}

	return r
}

// scale multiplies a point by a scalar
func (p point) scale(scalar float64) point {
	r := p
	for i := range r {
		r[i] *= scalar
	}

	return r
}

// testForallBut() is a helper function for the optimization to simplify a predicate
func testForallBut(f optfunc, xp point, sx simplex, h int) bool {
	for i := range sx {
		if i == h {
			continue
		} else {
			if f(xp) > f(sx[i]) {
				continue
			} else {
				return false
			}
		}
	}

	return true
}

// centroid calculates the centroid of a simplex of one dimensionality lower by omitting a point
func (s simplex) centroid(omit int) point {
	r := s[0]

	for i := range r {
		c := 0.0
		for j := range s {
			if j == omit {
				continue
			} else {
				c += s[j][i]
			}
		}

		r[i] = c / float64((len(s) - 1))
	}

	return r
}

// argMax finds the best point in the simplex for the optfunc
func (s simplex) argMax(f optfunc) (idx int) {
	v := f(s[0])
	idx = 0
	for i := 1; i < len(s); i++ {
		r := f(s[i])
		if r > v {
			v = r
			idx = i
		}
	}
	return
}

// argMin finds the worst point in the simplex for the optfunc
func (s simplex) argMin(f optfunc) (idx int) {
	v := f(s[0])
	idx = 0
	for i := 1; i < len(s); i++ {
		r := f(s[i])
		if r < v {
			v = r
			idx = i
		}
	}
	return
}

// stopCriterion tests if the stop criterion has been met
func stopCriterion(f optfunc, xs simplex) bool {
	r := make([]float64, len(xs))

	for i := range r {
		r[i] = f(xs[i])
	}

	sum := 0.0
	for _, v := range r {
		sum += v
	}
	avg := sum / float64(len(r))

	s := 0.0
	for _, v := range r {
		rs := (v - avg) * (v - avg)
		fmt.Printf("v is %v, result is %v\n", v, rs)
		s += rs
	}

	res := math.Sqrt(1 / (1.0 + float64(len(r))) * s)

	fmt.Printf("We are at %v toward the limit (%v) - s is %v, avg is %v\n", res, res < ε, s, avg)
	return (res > ε)
}
