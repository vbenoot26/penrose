package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var whiteImage = ebiten.NewImage(3, 3)

func init() {
	whiteImage.Fill(color.White)
}

// A file to keep track of all the game logic.
type Game struct {
	darts transSet
	kites transSet
}

func (g Game) Update() error { return nil }

func (g Game) Draw(screen *ebiten.Image) {
	// lets ignore gamestate for now and just draw one dart

	// coordinates to vertices TODO: make vertices the main data structure
	vertices := []ebiten.Vertex{}

	for _, co := range dart.points {
		newVertex := ebiten.Vertex{
			DstX:   getXCo(co),
			DstY:   getYCo(co),
			SrcX:   0,
			SrcY:   0,
			ColorR: 255,
			ColorG: 255,
			ColorB: 255,
			ColorA: 1,
		}

		vertices = append(vertices, newVertex)
	}

	// indices array
	indices := []uint16{0, 1, 2, 1, 2, 3}

	screen.DrawTriangles(vertices, indices, whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), &ebiten.DrawTrianglesOptions{})

}

func (g Game) Layout(outsidewidth, outsideHeigth int) (int, int) {
	return width, heigth
}

func getXCo(co coordinate) float32 {
	return float32((co.x * rescale) + (width / 2))
}

func getYCo(co coordinate) float32 {
	return float32((co.y * rescale) + (heigth / 2))
}
