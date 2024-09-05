package main

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	whiteImage = ebiten.NewImage(3, 3)

	emilieDartColor = color.NRGBA{
		R: 45,
		G: 41,
		B: 38,
		A: 255,
	}

	emilieKiteColor = color.NRGBA{
		R: 237,
		G: 111,
		B: 99,
		A: 255,
	}
)

const animationLength = 60

func init() {
	whiteImage.Fill(color.White)
}

// A file to keep track of all the game logic.
type Game struct {
	states           []state
	currentIteration int
	tick             int
}

func (g *Game) Update() error {
	if g.currentIteration >= maxIters {
		return nil
	}
	if g.tick >= animationLength-1 {
		g.currentIteration++
	}
	g.tick = (g.tick + 1) % animationLength
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.currentIteration > 0 {
		prevState := g.states[g.currentIteration-1]

		for trans := range prevState.dartTranses {
			drawDart(animationLength, trans, screen)
		}

		for trans := range prevState.kiteTranses {
			drawKite(animationLength, trans, screen)
		}
	}

	currentState := g.states[g.currentIteration]
	for trans := range currentState.dartTranses {
		drawDart(g.tick, trans, screen)
	}

	for trans := range currentState.kiteTranses {
		drawKite(g.tick, trans, screen)
	}
}

func drawDart(tick int, trans transformation, screen *ebiten.Image) {
	transformed := dart.getRescaledOnTick(tick).applyTransformation(trans)
	transformed.draw(dartColor, screen)
	transformed.drawBorder(screen)
}

func drawKite(tick int, trans transformation, screen *ebiten.Image) {
	transformed := kite.getRescaledOnTick(tick).applyTransformation(trans)
	transformed.draw(kiteColor, screen)
	transformed.drawBorder(screen)
}

// We presume that the reciever is either a variant of a kite or a dart. This is important for the way the triangles are drawn.
func (tile polygon) draw(color color.NRGBA, screen *ebiten.Image) {
	// coordinates to vertices
	vertices := []ebiten.Vertex{}

	for _, co := range tile {
		vertices = append(vertices, co.toVertex(color))
	}

	// indices array, since this is the same for kites and darts we just magic value hard code it.
	indices := []uint16{0, 1, 2, 0, 2, 3}

	screen.DrawTriangles(vertices, indices, whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), &ebiten.DrawTrianglesOptions{})
}

func (tile polygon) drawBorder(screen *ebiten.Image) {
	if !borders {
		return
	}

	for i := range tile {
		currentCo := tile[i]
		nextCo := tile[(i+1)%len(tile)]
		vector.StrokeLine(screen, getXCo(currentCo), getYCo(currentCo), getXCo(nextCo), getYCo(nextCo), 1, color.Black, true)
	}
}

func (poly *polygon) getRescaledOnTick(tick int) *polygon {
	animationStep := expSmooth(float64(tick) / animationLength)

	basicCos := poly

	resultcos := polygon{}
	for _, co := range *basicCos {
		resultcos = append(resultcos, co.scale(animationStep))
	}

	return &resultcos
}

func getSmoothingStep(tick int) float64 {
	return 1 - math.Pow(math.E, -float64(tick))
}

func linearSmooth(animationPart float64) float64 {
	return animationPart
}

func expSmooth(animationPart float64) float64 {
	speed := 10
	return 1 - math.Pow(math.E, -animationPart*float64(speed))
}

func sqrtSmooth(animationPart float64) float64 {
	return 1 - math.Sqrt(1-animationPart)
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
