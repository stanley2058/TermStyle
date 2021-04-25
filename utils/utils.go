package utils

import "math"

func hueToRgb(p float64, q float64, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	if t < 1/6 {
		return p + (q-p)*6*t
	}
	if t < 1/2 {
		return q
	}
	if t < 2/3 {
		return p + (q-p)*6*(2/3-t)
	}
	return p
}

// HslToRgb converted hsl to rgb (0<= h <360, 0<= s,l <=1)
func HslToRgb(h float64, s float64, l float64) (uint8, uint8, uint8) {
	h = math.Mod(h, 360)
	s = math.Min(math.Max(s, 0), 1)
	l = math.Min(math.Max(l, 0), 1)

	var rf, gf, bf, p, q float64

	if s == 0 {
		rf = l
		gf = l
		bf = l
	} else {
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p = 2*l - q

		rf = hueToRgb(p, q, h+1/3)
		gf = hueToRgb(p, q, h)
		bf = hueToRgb(p, q, h-1/3)
	}
	return uint8(rf * 255), uint8(gf * 255), uint8(bf * 255)
}
