# Experiments with ln

[ln](https://github.com/fogleman/ln) is a go library for creating vector 3d graphics.

Unlike a typical 3d renderer which outputs pixels, it outputs paths. 

I've been using it to make SVG files which I send to a pen plotter (the [Axidraw V3](https://axidraw.com).)

All you have to do to run these examples is have `go` and `ln` installed:

```
$ go get github.com/fogleman/ln/ln
```

and then you can just run the file.

```
go run 100cubes.go
```
## 100 cubes

![](100cubes.png?raw=true)

## Cube Tube

![](cubetube.png?raw=true)

## Pencilstorm

![](pencilstorm.png?raw=true)

## Random Walk

![](randomwalk.png?raw=true)


