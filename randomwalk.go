package main

import (
	"math/rand"
	"time"

	"github.com/fogleman/ln/ln"
)

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func addWalk(scene *ln.Scene, startAngle float64) {
	// The idea here is to take cubes and consistently rotate them around the Z axis
	// but for each step, randomly jitter the radius
	// and randomly increase the Z
	// keep doing that until the Z tops out at the "ceilng"

	cubesize := ln.Vector{0.25, 0.25, 0.25}
	origin := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}

	radius := 4.0
	z := 0.0
	cube := ln.NewCube(origin.Sub(cubesize), origin.Add(cubesize))
	angleInc := 360.0 / 30.0
	angle := startAngle

	zVel := 0.7
	radVel := -0.035

	for z < 40 {
		// adjust the speeds of change in radius and z
		// zVel += randFloat(-0.01, 0.1)
		// radVel += randFloat(-0.02, 0.02)

		// adjust the actual values
		radius += radVel
		z += zVel
		angle += angleInc

		center := ln.Vector{0, radius, z}
		m := ln.Translate(center).Rotate(up, ln.Radians(angle))
		scene.Add(ln.NewTransformedShape(cube, m))
	}

}

func main() {

	scene := ln.Scene{}
	up := ln.Vector{0, 0, 1}
	rand.Seed(time.Now().UnixNano())
	walks := 4
	for i := 0; i < walks; i++ {
		addWalk(&scene, (360.0/float64(walks))*float64(i))
	}
	scene.Add(ln.NewSphere(ln.Vector{0, 0, 45}, 2.0))

	// rendering parameters
	// 11
	width := 1325.0
	// 8.5
	height := 1024.0

	// camera parameters
	eye := ln.Vector{0, 0, 0}     // camera position
	center := ln.Vector{1, 1, 50} // camera looks at

	// rendering parameters
	fovy := 30.0 // vertical field of view in degrees
	znear := 0.1 // near z plane
	zfar := 50.0 // far z plane
	step := 0.01 // how finely to chop paths for visibility testing

	// compute the 2d paths for the 3d scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	paths.WriteToPNG("randomwalk.png", width, height)
	paths.WriteToSVG("randomwalk.svg", width, height)
}
