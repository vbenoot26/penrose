package main

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	whiteImage = ebiten.NewImage(3, 3)
	dartColor  = color.NRGBA{
		R: 238,
		G: 223,
		B: 226,
		A: 255,
	}
	// kiteColor = color.NRGBA{
	// 	R: 159,
	// 	G: 193,
	// 	B: 49,
	// 	A: 255,
	// }
	kiteColor = color.NRGBA{
		R: 159,
		G: 193,
		B: 49,
		A: 255,
	}
)

const animationLength = 60

func init() {
	whiteImage.Fill(color.White)
}

// A file to keep track of all the game logic.
type Game struct {
	darts transSet
	kites transSet
	tick  uint32
}

func (g *Game) Update() error {
	g.tick = uint32(math.Min(float64(g.tick+1), animationLength))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for trans := range g.darts {
		getDartCos(g.tick).applyTransformation(trans).draw(dartColor, screen)
	}

	for trans := range g.kites {
		getKiteCos(g.tick).applyTransformation(trans).draw(kiteColor, screen)
	}
}

// We presume that the reciever is either a variant of a kite or a dart. This is important for the way the triangles are drawn.
func (tile polygon) draw(color color.NRGBA, screen *ebiten.Image) {
	// coordinates to vertices
	vertices := []ebiten.Vertex{}

	for _, co := range tile.points {
		vertices = append(vertices, co.toVertex(color))
	}

	// indices array, since this is the same for kites and darts we just magic value hard code it.
	indices := []uint16{0, 1, 2, 0, 2, 3}

	screen.DrawTriangles(vertices, indices, whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), &ebiten.DrawTrianglesOptions{})
}

func getDartCos(tick uint32) *polygon {
	pointAngle := float64(tick) / float64(animationLength) * radian36

	return &polygon{
		[]coordinate{
			{0, 0},
			{math.Cos(pointAngle), math.Sin(pointAngle)},
			{1, 0},
			{math.Cos(pointAngle), -math.Sin(pointAngle)},
		},
	}
}

func getKiteCos(tick uint32) *polygon {
	pointAngle := float64(tick) / float64(animationLength) * radian36

	return &polygon{
		[]coordinate{
			{0, 0},
			{math.Cos(pointAngle), math.Sin(pointAngle)},
			{1 / math.Phi, 0},
			{math.Cos(pointAngle), -math.Sin(pointAngle)},
		},
	}
}

func (co *coordinate) toVertex(color color.NRGBA) ebiten.Vertex {
	return ebiten.Vertex{
		DstX:   getXCo(*co),
		DstY:   getYCo(*co),
		SrcX:   0,
		SrcY:   0,
		ColorR: float32(color.R) / 255,
		ColorG: float32(color.G) / 255,
		ColorB: float32(color.B) / 255,
		ColorA: float32(color.A) / 255,
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
