package main

import (
	"math"
	"os"

	"github.com/ajstarks/svgo"
)

var width = 1000
var heigth = 1000

func main() {
	canvas := svg.New(os.Stdout)
	canvas.Start(width, heigth)

	pointsx, pointsy := getDartCoords()

	canvas.Polygon(toCoordArray(pointsx), toCoordArray(pointsy), "fill:lightblue;stroke:black;stroke-width:2")

	canvas.End()
}

func getDartCoords() ([]float64, []float64) {
	radian36 := math.Pi / 5

	pointsx := []float64{
		0, math.Cos(radian36), 1, math.Cos(radian36),
	}

	pointsy := []float64{
		0, math.Sin(radian36), 0, -math.Sin(radian36),
	}

	return pointsx, pointsy
}

func toCoordArray(floatArr []float64) []int {
	result := []int{}
	for _, num := range floatArr {
		result = append(result, int((num*500)+500))
	}
	return result
}
