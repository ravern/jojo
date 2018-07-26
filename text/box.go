package text

import "github.com/ravernkoh/mojo"

// Box represents a box containing some text.
type Box struct {
	HorizontalAlignment Alignment
	VerticalAlignment   Alignment

	width  int
	height int

	content *StyledString
}

// Alignment represents the alignment of some text.
//
// Text needs to be aligned both in the horizontal and vertical direction. The
// alignment in both directions can be generalised as going from leading to
// trailing. The horizontal direction goes from left to right while the vertical
// direction goes from top to bottom.
type Alignment int

// Possible alignments of text.
const (
	AlignmentLeading Alignment = iota
	AlignmentCenter
	AlignmentTrailing
)

// NewBox creates a new box with the given content, width and height.
func NewBox(content string, width, height int) *Box {
	return nil
}

// Render renders some text.
func (c *Box) Render(s *mojo.Screen) error {
	return nil
}

// Click responds to a click.
func (c *Box) Click(x, y int) bool {
	return false
}

// Width returns the preferred width.
func (c *Box) Width() int {
	return 0
}

// Height returns the preferred height.
func (c *Box) Height() int {
	return 0
}

// SetWidth sets the preferred width of the box.
func (c *Box) SetWidth(width int) {
	c.width = width
}

// SetHeight sets the preferred height of the box.
func (c *Box) SetHeight(height int) {
	c.height = height
}
