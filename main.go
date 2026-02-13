package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth    = 800
	screenHeight   = 600
	minCellSize    = 2
	maxCellSize    = 50
	defaultCellSize = 10
)

type Cell struct {
	X, Y int
}

type Game struct {
	cells    map[Cell]bool
	paused   bool
	tickRate int
	tick     int
	offsetX  int
	offsetY  int
	cellSize int
}

func NewGame() *Game {
	g := &Game{
		paused:   true,
		tickRate: 10,
		cellSize: defaultCellSize,
	}
	Reset(g)

	return g
}

func Reset(g *Game) {
	g.cells = make(map[Cell]bool)
	g.cells[Cell{5, 5}] = true
	g.cells[Cell{6, 6}] = true
	g.cells[Cell{6, 7}] = true
	g.cells[Cell{5, 7}] = true
	g.cells[Cell{4, 7}] = true
}

func (g *Game) Update() error {
	// Toggle pause with Space
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = !g.paused
	}

	// Clear with C
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.cells = make(map[Cell]bool)
		g.paused = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		Reset(g)
	}

	// Zoom in/out with +/- keys
	if inpututil.IsKeyJustPressed(ebiten.KeyEqual) || inpututil.IsKeyJustPressed(ebiten.KeyKPAdd) {
		if g.cellSize < maxCellSize {
			g.cellSize++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyMinus) || inpututil.IsKeyJustPressed(ebiten.KeyKPSubtract) {
		if g.cellSize > minCellSize {
			g.cellSize--
		}
	}

	// Mouse input to toggle cells
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		cellX := (mx - g.offsetX) / g.cellSize
		cellY := (my - g.offsetY) / g.cellSize
		cell := Cell{cellX, cellY}

		if g.cells[cell] {
			delete(g.cells, cell)
		} else {
			g.cells[cell] = true
		}
	}

	// Update game state
	if !g.paused {
		g.tick++
		if g.tick >= g.tickRate {
			g.tick = 0
			g.step()
		}
	}

	return nil
}

func (g *Game) step() {
	// Count neighbors for all cells and their neighbors
	cellCount := len(g.cells)
	neighborCount := make(map[Cell]int, cellCount*8)

	for cell := range g.cells {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				neighbor := Cell{cell.X + dx, cell.Y + dy}
				neighborCount[neighbor]++
			}
		}
	}

	// Apply Game of Life rules
	newCells := make(map[Cell]bool, cellCount)

	for cell, count := range neighborCount {
		if count == 3 || (count == 2 && g.cells[cell]) {
			newCells[cell] = true
		}
	}

	g.cells = newCells
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 30, 255})

	// Draw grid
//	gridColor := color.RGBA{40, 40, 50, 255}
//	for x := 0; x < screenWidth; x += cellSize {
//		ebitenutil.DrawLine(screen, float64(x), 0, float64(x), screenHeight, gridColor)
//	}
//	for y := 0; y < screenHeight; y += cellSize {
//		ebitenutil.DrawLine(screen, 0, float64(y), screenWidth, float64(y), gridColor)
//	}

	// Draw cells as white rectangles
	cellColor := color.White
	for cell := range g.cells {
		x := cell.X*g.cellSize + g.offsetX
		y := cell.Y*g.cellSize + g.offsetY

		// Only draw if visible
		if x >= -g.cellSize && x < screenWidth && y >= -g.cellSize && y < screenHeight {
			ebitenutil.DrawRect(screen, float64(x), float64(y), float64(g.cellSize-1), float64(g.cellSize-1), cellColor)
		}
	}

	// Draw instructions
	status := "PAUSED"
	if !g.paused {
		status = "RUNNING"
	}
	ebitenutil.DebugPrint(screen, status+"\nSpace: Play/Pause | C: Clear | R: Reset | +/-: Zoom | Click: Toggle cells")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
