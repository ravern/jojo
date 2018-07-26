package mojo

// Text represents some text to be rendered.
type Text struct {
	HorizontalAlign TextAlign
	VerticalAlign   TextAlign

	text []rune
}

// TextAlign represents the alignment of some text.
//
// Text needs to be aligned both in the horizontal and vertical direction. The
// alignment in both directions can be generalised as going from leading to
// trailing. The horizontal direction goes from left to right while the vertical
// direction goes from top to bottom.
type TextAlign int

// Possible alignments of text.
const (
	TextAlignLeading TextAlign = iota
	TextAlignCenter
	TextAlignTrailing
)

// Render renders some text.
func (c *Text) Render(s *Screen) error {
	return nil
}

// Width returns the preferred width.
func (c *Text) Width() int {
	return 0
}

// Height returns the preferred height.
func (c *Text) Height() int {
	return 0
}
