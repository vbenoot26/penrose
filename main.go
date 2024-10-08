package main

import (
	"math"
	"os"

	svg "github.com/ajstarks/svgo"
)

var width = 1000
var heigth = 1000
var radian36 = math.Pi / 5

type coordinate struct {
	x float64
	y float64
}

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(width, heigth)

	canvas.Polygon(getXCoords(getKiteCoords()), getYCoords(getKiteCoords()), "fill:lightblue;stroke:black;stroke-width:2")

	canvas.End()
}

func getDartCoords() []coordinate {
	return []coordinate{
		{0, 0},
		{math.Cos(radian36), math.Sin(radian36)},
		{1, 0},
		{math.Cos(radian36), -math.Sin(radian36)},
	}
}

func getKiteCoords() []coordinate {
	return []coordinate{
		{0, 0},
		{math.Cos(radian36), math.Sin(radian36)},
		{1 / math.Phi, 0},
		{math.Cos(radian36), -math.Sin(radian36)},
	}
}

func getXCoords(coordinates []coordinate) []int {
	result := []int{}
	for _, coord := range coordinates {
		x := coord.x
		result = append(result, int((x*500)+500))
	}
	return result
}

func getYCoords(coordinates []coordinate) []int {
	result := []int{}
	for _, coord := range coordinates {
		y := coord.y
		result = append(result, int((y*500)+500))
	}
	return result
}
