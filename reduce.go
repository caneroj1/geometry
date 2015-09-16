package main

import "math"

type vector struct {
	x float64
	y float64
}

// Reduce modifies the points in the slice through
// the Ramos-Douglas-Peucker algorithm.
func Reduce(p Points, eps float64) Points {
	if len(p) == 1 {
		return p
	}

	dmax := float64(0)
	idx := 0
	end := len(p) - 1

	for i := 1; i < end-1; i++ {
		d := shortestDistanceToSegment(p[i], p[0], p[end])
		if d > dmax {
			dmax = d
			idx = i
		}
	}

	if dmax > eps {
		l := Reduce(p[0:idx+1], eps)
		r := Reduce(p[idx:end+1], eps)
		return append(l[:len(l)-1], r...)
	}

	return Points{p[0], p[end]}
}

func shortestDistanceToSegment(p, s, e Point) float64 {
	v1 := makeVector(s, e)
	m1 := v1.magnitude()
	num := float64(v1.y*p.X - v1.x*p.Y + e.X*s.Y - e.Y*s.X)
	return math.Abs(num) / m1
}

func makeVector(p1, p2 Point) vector {
	return vector{
		p2.X - p1.X,
		p2.Y - p1.Y,
	}
}

func (v vector) magnitude() float64 {
	xs := float64(v.x * v.x)
	ys := float64(v.y * v.y)
	return math.Sqrt(xs + ys)
}
