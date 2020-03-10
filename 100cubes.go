package main

import "github.com/fogleman/ln/ln"

func main() {
	// create a scene and add a cube

	scene := ln.Scene{}
	cubesize := ln.Vector{0.25, 0.25, 0.25}
	origin := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	cube := ln.NewCube(origin.Sub(cubesize), origin.Add(cubesize))

	// todo figure out how to make a ring of circles

	for x := 1.0; x <= 10; x++ {
		for y := 1.0; y <= 10; y++ {
			center := ln.Vector{x, y, (x + y) * .3}
			m := ln.Rotate(up, ln.Radians((x+y)*10)).Translate(center)
			scene.Add(ln.NewTransformedShape(cube, m))
		}
	}
	// scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

	// rendering parameters
	width := 1024.0
	height := 1024.0

	// camera parameters
	eye := ln.Vector{-5, -5, 3}    // camera position
	center := ln.Vector{20, 20, 0} // camera looks at

	// rendering parameters
	fovy := 50.0 // vertical field of view in degrees
	znear := 0.1 // near z plane
	zfar := 50.0 // far z plane
	step := 0.01 // how finely to chop paths for visibility testing

	// compute the 2d paths for the 3d scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	paths.WriteToPNG("100cubes.png", width, height)
	paths.WriteToSVG("100cubes.svg", width, height)
}
