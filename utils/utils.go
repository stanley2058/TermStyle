package utils

import "math"

// HslToRgb converted hsl to rgb (0<= h <360, 0<= s,l <=1)
func HslToRgb(h float64, s float64, l float64) (r uint8, g uint8, b uint8) {
	h = math.Mod(h, 360)
	s = math.Min(math.Max(s, 0), 1)
	l = math.Min(math.Max(l, 0), 1)

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var rp, gp, bp float64
	if h < 60 {
		rp = c
		gp = x
		bp = 0
	} else if h < 120 {
		rp = x
		gp = c
		bp = 0
	} else if h < 180 {
		rp = 0
		gp = c
		bp = x
	} else if h < 240 {
		rp = 0
		gp = x
		bp = c
	} else if h < 300 {
		rp = x
		gp = 0
		bp = c
	} else if h < 360 {
		rp = c
		gp = 0
		bp = x
	}

	r = uint8((rp + m) * 255)
	g = uint8((gp + m) * 255)
	b = uint8((bp + m) * 255)
	return r, g, b
}
