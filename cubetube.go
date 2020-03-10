package main

import (
	"math/rand"
	"time"

	"github.com/fogleman/ln/ln"
)

func main() {
	// create a scene and add a cube

	scene := ln.Scene{}
	cubesize := ln.Vector{0.25, 0.25, 0.25}
	origin := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}

	rand.Seed(time.Now().UnixNano())

	for row := 0.0; row <= 24.0; row += 1.0 {
		// 50 50 chance of just not drawing a given row
		if rand.Intn(2) == 0 {
			continue
		}
		var shape ln.Shape
		switch rand.Intn(1) {
		case 0:
			shape = ln.NewCube(origin.Sub(cubesize), origin.Add(cubesize))
		case 1:
			shape = ln.NewSphere(origin, 0.5)
		}

		// randomly select the number of objects
		objects := 25 + rand.Intn(22)
		for angle := 0.0; angle < 360.0; angle += 360.0 / float64(objects) {
			spacing := row
			if row > 10.0 {
				spacing = 24.0 - row
			}
			center := ln.Vector{0, 5.0 + spacing, row}
			m := ln.Translate(center).Rotate(up, ln.Radians(angle+(row*3)))
			scene.Add(ln.NewTransformedShape(shape, m))
		}
	}

	// scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

	// rendering parameters
	// 11
	width := 1325.0
	// 8.5
	height := 1024.0

	// camera parameters
	eye := ln.Vector{3, 3, 0}     // camera position
	center := ln.Vector{5, 5, 50} // camera looks at

	// rendering parameters
	fovy := 50.0 // vertical field of view in degrees
	znear := 0.1 // near z plane
	zfar := 50.0 // far z plane
	step := 0.01 // how finely to chop paths for visibility testing

	// compute the 2d paths for the 3d scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	paths.WriteToPNG("cubetube.png", width, height)
	paths.WriteToSVG("cubetube.svg", width, height)
}
