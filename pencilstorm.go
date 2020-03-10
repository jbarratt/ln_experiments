package main

import (
	"math/rand"
	"time"

	"github.com/fogleman/ln/ln"
)

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main() {

	scene := ln.Scene{}
	// origin := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	right := ln.Vector{0, 1, 0}

	rand.Seed(time.Now().UnixNano())

	mesh, err := ln.LoadOBJ("pencil.obj")
	if err != nil {
		panic(err)
	}

	mesh.UnitCube()

	for i := 0; i < 100; i++ {
		center := ln.Vector{randFloat(0.5, 5.0), randFloat(0.5, 5.0), randFloat(1.0, 4.0)}
		m := ln.Rotate(up, ln.Radians(randFloat(0.0, 360.0))).Rotate(right, ln.Radians(randFloat(0.0, 360.0))).Translate(center)
		// m := ln.Translate(center)
		scene.Add(ln.NewTransformedShape(mesh, m))
	}

	// rendering parameters
	// 11
	width := 1325.0
	// 8.5
	height := 1024.0

	// camera parameters
	eye := ln.Vector{0, 0, 2}        // camera position
	lookat := ln.Vector{10, 10, 2.5} // camera looks at

	// rendering parameters
	fovy := 60.0 // vertical field of view in degrees
	znear := 0.1 // near z plane
	zfar := 50.0 // far z plane
	step := 0.01 // how finely to chop paths for visibility testing

	// compute the 2d paths for the 3d scene
	paths := scene.Render(eye, lookat, up, width, height, fovy, znear, zfar, step)

	paths.WriteToPNG("pencilstorm.png", width, height)
	paths.WriteToSVG("pencilstorm.svg", width, height)
}
