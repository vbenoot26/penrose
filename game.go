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

func (g *Game) Update() error { return nil }

func (g *Game) Draw(screen *ebiten.Image) {
	dartTranses, kiteTranses := dartReplace(idTransform)

	for _, trans := range dartTranses {
		drawDart(trans, screen)
	}

	for _, trans := range kiteTranses {
		drawKite(trans, screen)
	}
}

func drawDart(trans transformation, screen *ebiten.Image) {
	// coordinates to vertices
	vertices := []ebiten.Vertex{}

	for _, co := range dart.applyTransformation(trans).points {
		vertices = append(vertices, newDartVertex(co))
	}

	// indices array
	indices := []uint16{0, 1, 2, 0, 2, 3}

	screen.DrawTriangles(vertices, indices, whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), &ebiten.DrawTrianglesOptions{})
}

func drawKite(trans transformation, screen *ebiten.Image) {
	// coordinates to vertices
	vertices := []ebiten.Vertex{}

	for _, co := range kite.applyTransformation(trans).points {
		vertices = append(vertices, newKiteVertex(co))
	}

	// indices array
	indices := []uint16{0, 1, 2, 0, 2, 3}

	screen.DrawTriangles(vertices, indices, whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), &ebiten.DrawTrianglesOptions{})
}

func newDartVertex(co coordinate) ebiten.Vertex {
	return ebiten.Vertex{
		DstX:   getXCo(co),
		DstY:   getYCo(co),
		SrcX:   0,
		SrcY:   0,
		ColorR: 0.9333333333333333,
		ColorG: 0.8745098039215686,
		ColorB: 0.8862745098039215,
		ColorA: 1,
	}
}

func newKiteVertex(co coordinate) ebiten.Vertex {
	return ebiten.Vertex{
		DstX:   getXCo(co),
		DstY:   getYCo(co),
		SrcX:   0,
		SrcY:   0,
		ColorR: 0.6235294117647059,
		ColorG: 0.7568627450980392,
		ColorB: 0.19215686274509805,
		ColorA: 1,
	}
}

func (g *Game) Layout(outsidewidth, outsideHeigth int) (int, int) {
	return width, heigth
}

func getXCo(co coordinate) float32 {
	return float32((co.x * rescale) + (width / 2))
}

func getYCo(co coordinate) float32 {
	return float32((co.y * rescale) + (heigth / 2))
}
