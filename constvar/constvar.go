package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	area := PI * m.Pow(raio, 2)
	fmt.Println(area)
	fmt.Println(a, b, c, d)
	println("ué")
}
