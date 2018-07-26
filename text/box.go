package text

import "github.com/ravernkoh/mojo"

// Box represents a box containing some text.
type Box struct {
	HorizontalAlignment Alignment
	VerticalAlignment   Alignment

	width  int
	height int

	text *StyledString
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

// NewBox creates a new box with the given text, width and height.
func NewBox(text string, width, height int) *Box {
	return nil
}

// NewBoxStyled creates a new box with the given text with styles, width and
// height.
func NewBoxStyled(text *StyledString, width, height int) *Box {
	return nil
}

// Render renders the text onto the given screen.
func (c *Box) Render(s *mojo.Screen, e mojo.Event) error {
	return nil
}

// Text returns the text.
func (c *Box) Text() string {
	return c.text.String()
}

// SetText sets the text.
func (c *Box) SetText(text string) {
	c.text = NewStyledString(text)
}

// TextStyled returns the text with styles.
func (c *Box) TextStyled() *StyledString {
	return c.text
}

// SetTextStyled sets the text with styles.
func (c *Box) SetTextStyled(text *StyledString) {
	c.text = text
}

// Width returns the preferred width.
func (c *Box) Width() int {
	return 0
}

// SetWidth sets the preferred width of the box.
func (c *Box) SetWidth(width int) {
	c.width = width
}

// Height returns the preferred height.
func (c *Box) Height() int {
	return 0
}

// SetHeight sets the preferred height of the box.
func (c *Box) SetHeight(height int) {
	c.height = height
}
