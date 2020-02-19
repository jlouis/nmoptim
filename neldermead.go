package nmoptim

import (
	"log"
	"math"
)

const (
	max = 1000     // arbitrarily chosen value for now
	ε   = 0.000001 // Stopping criterion point
	α   = 1.0
	β   = 0.5
	γ   = 2.0
)

// point is the type of points in ℝ^n
type point []float64

// simplex is the type used to represent a simplex
type simplex []point

// Optimize function f with Nelder-Mead. start points to a slice of starting points
// It is the responsibility of the caller to make sure the dimensionality is correct.
//
// The cf function is a constraint function which can be used to clamp the space in which
// your optimization runs. This avoid the algorithm picking extreme points, ouside the
// space you want to operate on.
func Optimize(f func([]float64) float64, start [][]float64, cf func([]float64)) ([]float64, int, int) {
	evaluations := 0
	n := len(start)
	c := len(start[0])
	points := make([]point, 0)
	fv := make([]float64, n)

	for _, p := range start {
		points = append(points, point(p))
	}
	sx := simplex(points)
	if n != c+1 {
		log.Fatalf("Can't optimize with too few starting points. D:%v, SP:%v", n, c)
	}

	// Set up initial values
	for i := range fv {
		if cf != nil {
			cf(sx[i])
		}
		fv[i] = f(sx[i])
		evaluations++
	}

	k := 0
	for ; k < max; k++ {
		// Find the largest index
		vg := 0
		for i := range fv {
			if fv[i] > fv[vg] {
				vg = i
			}
		}

		// Find the smallest index
		vs := 0
		for i := range fv {
			if fv[i] < fv[vs] {
				vs = i
			}
		}

		// Second largest index
		vh := vs
		for i := range fv {
			if fv[i] > fv[vh] && fv[i] < fv[vg] {
				vh = i
			}
		}

		vm := sx.centroid(vg)

		vr := add(vm, sub(vm, sx[vg]).scale(α))
		if cf != nil {
			cf(vr)
		}
		fr := f(vr)
		evaluations++

		if fr < fv[vh] && fr >= fv[vs] {
			// Replace
			fv[vg] = fr
			sx[vg] = vr
		}

		// Investigate a step further
		if fr < fv[vs] {
			ve := add(vm, sub(vr, vm).scale(γ))
			if cf != nil {
				cf(ve)
			}

			fe := f(ve)
			evaluations++

			if fe < fr {
				sx[vg] = ve
				fv[vg] = fe
			} else {
				sx[vg] = vr
				fv[vg] = fr
			}
		}

		// Check contraction
		if fr >= fv[vh] {
			var vc point
			var fc float64
			if fr < fv[vg] && fr >= fv[vh] {
				// Outside contraction
				vc = add(vm, sub(vr, vm).scale(β))
			} else {
				// Inside contraction
				vc = sub(vm, sub(vm, sx[vg]).scale(β))
			}

			if cf != nil {
				cf(vc)
			}
			fc = f(vc)
			evaluations++

			if fc < fv[vg] {
				sx[vg] = vc
				fv[vg] = fc
			} else {
				for i := range sx {
					if i != vs {
						sx[i] = add(sx[vs], sub(sx[i], sx[vs]).scale(0.5))
					}
				}

				if cf != nil {
					cf(sx[vg])
				}
				fv[vg] = f(sx[vg])
				evaluations++

				if cf != nil {
					cf(sx[vh])
				}
				fv[vh] = f(sx[vh])
				evaluations++
			}
		}

		fsum := 0.0
		for _, v := range fv {
			fsum += v
		}

		favg := fsum / float64(len(fv))

		s := 0.0
		for _, v := range fv {
			s += math.Pow(v-favg, 2.0)
		}

		s = s * (1.0 / (float64(len(fv)) + 1.0))
		s = math.Sqrt(s)
		if s < ε {
			break
		}
	}

	vs := 0
	for i := range fv {
		if fv[i] < fv[vs] {
			vs = i
		}
	}

	return sx[vs], k, evaluations
}

// sub perform point subtraction
func sub(x point, y point) point {
	r := make(point, len(x))

	for i := range y {
		r[i] = x[i] - y[i]
	}

	return r
}

// add perform point addition
func add(x point, y point) point {
	r := make(point, len(x))

	for i := range y {
		r[i] = x[i] + y[i]
	}

	return r
}

// scale multiplies a point by a scalar
func (p point) scale(scalar float64) point {
	r := make(point, len(p))

	for i := range r {
		r[i] = scalar * p[i]
	}

	return r
}

// centroid calculates the centroid of a simplex of one dimensionality lower by omitting a point
func (s simplex) centroid(omit int) point {
	r := make(point, len(s[0]))

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
