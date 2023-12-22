package main

import "math"

func Perimeter(c string, d float64) float64 {
	if c == "c" {
		return 2 * math.Pi * d
	}
	return d * d
}
