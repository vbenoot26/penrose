package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const width = 1500
const heigth = 1000
const rescale = 400

const radian36 = math.Pi / 5
const scaleFactor = math.Phi - 1.0

var dart = polygon{
	[]coordinate{
		{0, 0},
		{math.Cos(radian36), math.Sin(radian36)},
		{1, 0},
		{math.Cos(radian36), -math.Sin(radian36)},
	},
}

var kite = polygon{
	[]coordinate{
		{0, 0},
		{math.Cos(radian36), math.Sin(radian36)},
		{1 / math.Phi, 0},
		{math.Cos(radian36), -math.Sin(radian36)},
	},
}

type coordinate struct {
	x float64
	y float64
}

type transformation struct {
	amountOfRotation int // Since all rotation are multiples of radian36, we just save the factor in front of radian36
	translation      coordinate
	rescales         int // rescales will always happen with the same factor: see scalefactor
}

type polygon struct {
	points []coordinate
}

func main() {
	ebiten.SetWindowSize(width, heigth)
	ebiten.SetWindowTitle("Penrose")
	log.Print("Initted")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func getXCoords(coordinates []coordinate) []int {
	result := []int{}
	for _, coord := range coordinates {
		x := coord.x
		result = append(result, int((x*rescale)+(width/2)))
	}
	return result
}

func getYCoords(coordinates []coordinate) []int {
	result := []int{}
	for _, coord := range coordinates {
		y := coord.y
		result = append(result, int((y*rescale)+(heigth/2)))
	}
	return result
}
