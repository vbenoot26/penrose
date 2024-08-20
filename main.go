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
var scaleFactor = 2 - math.Phi

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
	rotation    float64
	translation coordinate
	rescales    int // rescales will always happen with the same factor: see scalefactor
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

	for _, shape := range kiteReplace() {
		drawPolygon(canvas, shape)
	}
}

func drawPolygon(canvas *svg.SVG, shape polygon) {
	canvas.Polygon(getXCoords(shape.points), getYCoords(shape.points), "fill:lightblue;stroke:black;stroke-width:2")
}

func (shape *polygon) applyTransformation(transform transformation) polygon {
	newPoints := []coordinate{}
	for _, coord := range shape.points {
		newPoint := applyTransformation(coord, transform)
		newPoints = append(newPoints, newPoint)
	}
	return polygon{newPoints}
}

func applyTransformation(coord coordinate, transform transformation) coordinate {
	// rescales
	rescale := math.Pow(scaleFactor, float64(transform.rescales))
	newCoord := coordinate{
		coord.x * rescale,
		coord.y * rescale,
	}

	// rotation
	angle := transform.rotation
	rotationMatrix := [][]float64{
		{math.Cos(angle), -math.Sin(angle)},
		{math.Sin(angle), math.Cos(angle)},
	}

	newCoord = matrixTransform(coord, rotationMatrix)

	// translation
	newCoord = coordinate{
		newCoord.x + transform.translation.x,
		newCoord.y + transform.translation.y,
	}

	return newCoord
}

func matrixTransform(coord coordinate, matrix [][]float64) coordinate {
	return coordinate{
		coord.x*matrix[0][0] + coord.y*matrix[0][1],
		coord.x*matrix[1][0] + coord.y*matrix[1][1],
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

func kiteReplace() []polygon {
	return []polygon{
		dart.applyTransformation(transformation{
			0, coordinate{0, 0}, 1,
		}),
		kite.applyTransformation(transformation{
			6 * math.Pi / 5,
			coordinate{math.Cos(radian36), math.Sin(radian36)},
			1,
		}),
		kite.applyTransformation(transformation{
			-6 * math.Pi / 5,
			coordinate{math.Cos(radian36), -math.Sin(radian36)},
			1,
		}),
	}
}
