# Conway's Game of Life in Go

A simple implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) using Go and the Ebitengine game engine.

## About

This project was written by Claude because the author is lazy - one of [Larry Wall's three virtues of a programmer](http://threevirtues.com/). It serves as a learning exercise for:

- **Go programming language** - Understanding Go's syntax, structs, and methods
- **Ebitengine (Ebiten)** - A simple 2D game engine for Go

## The Game of Life

Conway's Game of Life is a cellular automaton where cells live or die based on simple rules:

1. Any live cell with 2-3 live neighbors survives
2. Any dead cell with exactly 3 live neighbors becomes alive
3. All other cells die or stay dead

Despite these simple rules, complex patterns emerge!

## Controls

### Simulation
- **Space** - Play/Pause the simulation
- **Left Click** - Toggle cells on/off (draw/erase)

### Grid Management
- **C** - Clear all cells
- **F** - Fill all visible cells
- **R** - Fill 75% of cells randomly

### View
- **+/=** - Zoom in (increase cell size)
- **-** - Zoom out (decrease cell size)
- **G** - Toggle grid lines
- **D** - Toggle debug stats (FPS, TPS, cell count)

## Running

```bash
./gameoflife
```

Or build from source:

```bash
go build -o gameoflife
```

## Testing

Run the comprehensive test suite:

```bash
go test -v
```

Tests verify correct Game of Life rules including:
- Survival with 2-3 neighbors
- Death by underpopulation/overpopulation
- Reproduction with exactly 3 neighbors
- Classic patterns (Blinker, Block, Glider)

## Ebitengine Interface

Ebitengine requires your game struct to implement the `ebiten.Game` interface with three methods:

### 1. `Update() error`
Called every tick (typically 60 FPS) for game logic and input handling.

```go
func (g *Game) Update() error {
    // Handle input, update game state
    return nil
}
```

### 2. `Draw(screen *ebiten.Image)`
Called every frame to render graphics to the screen.

```go
func (g *Game) Draw(screen *ebiten.Image) {
    // Draw everything here
}
```

### 3. `Layout(outsideWidth, outsideHeight int) (int, int)`
Returns the logical screen dimensions for the coordinate system.

```go
func (g *Game) Layout(w, h int) (int, int) {
    return screenWidth, screenHeight
}
```

Go uses **implicit interface satisfaction** - any struct with these three methods automatically implements `ebiten.Game`. No explicit inheritance or declarations needed!

## Features

- **Infinite grid** - Sparse map representation only stores living cells
- **Configurable zoom** - Adjust cell size from 2px to 50px to fit thousands of cells
- **Random generation** - Quickly generate interesting patterns
- **Performance optimized** - Map pre-allocation and efficient rectangle rendering
- **Debug mode** - Monitor FPS, TPS, and cell count in real-time

## Implementation Details

- Uses a sparse representation (map) for the infinite grid - only stores living cells
- Cells rendered as white rectangles for maximum performance
- Grid-based coordinate system with dynamic cell sizing (2-50 pixels)
- Efficient neighbor counting using pre-allocated maps
- Optimized drawing with visibility culling

## License

Do whatever you want with it. But don't vote mean people, please.
