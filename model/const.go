package model

import (
	"image/color"
)

const (
	// BlockSize is the size of the labyrinth unit in pixels.
	BlockSize = 40

	// Rows is the number of rows in the labyrinth
	Rows = 33
	// Cols is the number of columns in the labyrinth
	Cols = 33

	// LabWidth is the width of the labyrinth's image in pixels.
	LabWidth = Cols * BlockSize

	// LabHeight is the height of the labyrinth's image in pixels.
	LabHeight = Rows * BlockSize
)

// V is the moving speed of gopher in pixel/sec.
const V = BlockSize * 2.0
 

// Type of the unit of the labyrinth
type Block int

// Block types of the labyrinth
const (
	// Empty block (free passage)
	BlockEmpty Block = iota
	// Wall block
	BlockWall
)

// Color "constants"
var (
	Black   = color.RGBA{A: 0xff}
	WallCol = color.RGBA{0xe0, 0xe0, 0xe0, 0xff}
)