package main

import (
	"testing"
)

func TestGameOfLifeRules(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[Cell]bool
		expected map[Cell]bool
	}{
		{
			name: "Cell with 2 neighbors survives",
			initial: map[Cell]bool{
				{0, 0}: true,
				{1, 0}: true,
				{0, 1}: true,
			},
			expected: map[Cell]bool{
				{0, 0}: true,
				{1, 0}: true,
				{0, 1}: true,
				{1, 1}: true,
			},
		},
		{
			name: "Cell with 3 neighbors survives",
			initial: map[Cell]bool{
				{0, 0}: true,
				{1, 0}: true,
				{0, 1}: true,
				{1, 1}: true,
			},
			expected: map[Cell]bool{
				{0, 0}: true,
				{1, 0}: true,
				{0, 1}: true,
				{1, 1}: true,
			},
		},
		{
			name: "Cell with fewer than 2 neighbors dies (underpopulation)",
			initial: map[Cell]bool{
				{0, 0}: true,
				{5, 5}: true,
			},
			expected: map[Cell]bool{},
		},
		{
			name: "Cell with more than 3 neighbors dies (overpopulation)",
			initial: map[Cell]bool{
				{1, 1}: true,
				{0, 0}: true,
				{1, 0}: true,
				{2, 0}: true,
				{0, 1}: true,
				{2, 1}: true,
			},
			expected: map[Cell]bool{
				{0, 0}:  true,
				{2, 0}:  true,
				{0, 1}:  true,
				{2, 1}:  true,
				{1, -1}: true,
				{1, 2}:  true,
			},
		},
		{
			name: "Dead cell with exactly 3 neighbors becomes alive (reproduction)",
			initial: map[Cell]bool{
				{0, 0}: true,
				{1, 0}: true,
				{2, 0}: true,
			},
			expected: map[Cell]bool{
				{1, 0}:  true,
				{1, 1}:  true,
				{1, -1}: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := &Game{
				cells:    tt.initial,
				paused:   false,
				tickRate: 1,
			}

			game.step()

			if len(game.cells) != len(tt.expected) {
				t.Errorf("Expected %d cells, got %d", len(tt.expected), len(game.cells))
			}

			for cell := range tt.expected {
				if !game.cells[cell] {
					t.Errorf("Expected cell %v to be alive", cell)
				}
			}

			for cell := range game.cells {
				if !tt.expected[cell] {
					t.Errorf("Expected cell %v to be dead", cell)
				}
			}
		})
	}
}

func TestBlinker(t *testing.T) {
	game := &Game{
		cells: map[Cell]bool{
			{1, 0}: true,
			{1, 1}: true,
			{1, 2}: true,
		},
		paused:   false,
		tickRate: 1,
	}

	expectedPhase1 := map[Cell]bool{
		{0, 1}: true,
		{1, 1}: true,
		{2, 1}: true,
	}

	game.step()

	if len(game.cells) != len(expectedPhase1) {
		t.Errorf("Blinker phase 1: Expected %d cells, got %d", len(expectedPhase1), len(game.cells))
	}

	for cell := range expectedPhase1 {
		if !game.cells[cell] {
			t.Errorf("Blinker phase 1: Expected cell %v to be alive", cell)
		}
	}

	expectedPhase2 := map[Cell]bool{
		{1, 0}: true,
		{1, 1}: true,
		{1, 2}: true,
	}

	game.step()

	if len(game.cells) != len(expectedPhase2) {
		t.Errorf("Blinker phase 2: Expected %d cells, got %d", len(expectedPhase2), len(game.cells))
	}

	for cell := range expectedPhase2 {
		if !game.cells[cell] {
			t.Errorf("Blinker phase 2: Expected cell %v to be alive", cell)
		}
	}
}

func TestBlock(t *testing.T) {
	initial := map[Cell]bool{
		{0, 0}: true,
		{1, 0}: true,
		{0, 1}: true,
		{1, 1}: true,
	}

	game := &Game{
		cells:    copyMap(initial),
		paused:   false,
		tickRate: 1,
	}

	game.step()

	if len(game.cells) != len(initial) {
		t.Errorf("Block (still life): Expected %d cells, got %d", len(initial), len(game.cells))
	}

	for cell := range initial {
		if !game.cells[cell] {
			t.Errorf("Block (still life): Expected cell %v to be alive", cell)
		}
	}
}

func TestGlider(t *testing.T) {
	game := &Game{
		cells: map[Cell]bool{
			{1, 0}: true,
			{2, 1}: true,
			{0, 2}: true,
			{1, 2}: true,
			{2, 2}: true,
		},
		paused:   false,
		tickRate: 1,
	}

	initialCount := len(game.cells)

	for i := 0; i < 4; i++ {
		game.step()
	}

	if len(game.cells) != initialCount {
		t.Errorf("Glider: Expected to maintain %d cells after 4 steps, got %d", initialCount, len(game.cells))
	}

	minX, maxX := 1000, -1000
	minY, maxY := 1000, -1000
	for cell := range game.cells {
		if cell.X < minX {
			minX = cell.X
		}
		if cell.X > maxX {
			maxX = cell.X
		}
		if cell.Y < minY {
			minY = cell.Y
		}
		if cell.Y > maxY {
			maxY = cell.Y
		}
	}

	if maxX <= 2 || maxY <= 2 {
		t.Errorf("Glider: Expected glider to move (max position: x=%d, y=%d)", maxX, maxY)
	}
}

func copyMap(m map[Cell]bool) map[Cell]bool {
	result := make(map[Cell]bool)
	for k, v := range m {
		result[k] = v
	}
	return result
}
