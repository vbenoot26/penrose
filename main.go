package main

import (
	"fmt"
	"math"
	"net/http"

	svg "github.com/ajstarks/svgo"
)

var width = 1000
var heigth = 1000
var radian36 = math.Pi / 5
var scaleFactor = math.Phi - 1.0

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
	http.HandleFunc("/", draw)
	fmt.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func draw(writer http.ResponseWriter, _ *http.Request) {
	canvas := svg.New(writer)
	canvas.Start(width, heigth)
	defer canvas.End()

	for _, shape := range drawPolygons() {
		drawPolygon(canvas, shape)
	}
}

func drawPolygon(canvas *svg.SVG, shape polygon) {
	canvas.Polygon(getXCoords(shape.points), getYCoords(shape.points), "fill:lightblue;stroke:black;stroke-width:2")
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
