package main

import (
	"flag"
	"image/color"
	"log"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"gopkg.in/go-playground/colors.v1"
)

const (
	width   = 1500
	heigth  = 1000
	rescale = 400

	radian36    = math.Pi / 5
	scaleFactor = math.Phi - 1.0
)

var (
	dart = polygon{
		{0, 0},
		{math.Cos(radian36), math.Sin(radian36)},
		{1, 0},
		{math.Cos(radian36), -math.Sin(radian36)},
	}

	kite = polygon{
		{0, 0},
		{math.Cos(radian36), math.Sin(radian36)},
		{1 / math.Phi, 0},
		{math.Cos(radian36), -math.Sin(radian36)},
	}

	idTransform = transformation{
		amountOfRotation: 0,
		translation:      coordinate{0, 0},
		rescales:         0,
	}

	// Options
	maxIters int
	borders  bool

	dartColor color.NRGBA
	kiteColor color.NRGBA
)

type coordinate struct {
	x float64
	y float64
}

type transformation struct {
	amountOfRotation int // Since all rotation are multiples of radian36, we just save the factor in front of radian36
	translation      coordinate
	rescales         int // rescales will always happen with the same factor: see scalefactor
}

type polygon []coordinate

type state struct {
	dartTranses transSet
	kiteTranses transSet
}

func main() {
	initArgs()
	ebiten.SetWindowSize(width, heigth)
	ebiten.SetWindowTitle("Penrose")
	game := Game{
		calculateDrawings(), 0, 0,
	}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func initArgs() {
	var kiteHex, dartHex string

	flag.BoolVar(&borders, "borders", false, "Turns borders on")
	flag.StringVar(&dartHex, "dartColor", "#eedfe2", "Define the color of a dart in hex")
	flag.StringVar(&kiteHex, "kiteColor", "#9fc131", "Define the color of a kite in hex")
	flag.Parse()

	dartColor = getColor(dartHex)
	kiteColor = getColor(kiteHex)

	itersInput, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal("Something went wrong converting the input to an int")
	}
	maxIters = itersInput
}

func getColor(hexStr string) color.NRGBA {
	hexCol, err := colors.ParseHEX(hexStr)
	if err != nil {
		log.Fatal(err)
	}

	rgbColors := hexCol.ToRGBA()
	return color.NRGBA{
		R: rgbColors.R,
		B: rgbColors.B,
		G: rgbColors.G,
		A: 255,
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
