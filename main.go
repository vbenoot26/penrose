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

	for i := 0; i < 5; i++ {
		rotaionTransform := transformation{
			2.0 * radian36 * float64(i),
			coordinate{0, 0},
		}

		drawPolygon(canvas, dart.applyTransformation(rotaionTransform))
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
	// rotation
	angle := transform.rotation
	rotationMatrix := [][]float64{
		{math.Cos(angle), -math.Sin(angle)},
		{math.Sin(angle), math.Cos(angle)},
	}
	// TODO: translation

	return matrixTransform(coord, rotationMatrix)
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
