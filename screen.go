package mojo

// Screen represents a collection of cells to be rendered to.
type Screen struct {
	width  int
	height int
	cells  [][]Cell
}

// Cell represents a single cell to be rendered to.
type Cell struct {
}

// NewScreen creates a new screen with the given width and height.
func NewScreen(width, height int) *Screen {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}

	return &Screen{
		width:  width,
		height: height,
		cells:  cells,
	}
}

// Subscreen returns a section at the given x and y with the given width and
// height.
//
// Rendering to the subscreen will also change the original screen, as they
// both point to the same collection of cells internally.
//
// It is assumed that the coordinates are valid.
func (s *Screen) Subscreen(x, y, width, height int) *Screen {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = s.cells[y+i][x : x+width]
	}

	return &Screen{
		width:  width,
		height: height,
		cells:  cells,
	}
}

// Width returns the width.
func (s *Screen) Width() int {
	return s.width
}

// Height returns the height.
func (s *Screen) Height() int {
	return s.height
}
