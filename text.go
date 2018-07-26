package mojo

// Text represents a box containing some text.
type Text struct {
	HorizontalTextAlign TextAlign
	VerticalTextAlign   TextAlign

	width  int
	height int

	text *StyledString
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

// NewText creates a new box with the given text, width and height.
func NewText(text string, width, height int) *Text {
	return nil
}

// NewTextStyled creates a new box with the given text with styles, width and
// height.
func NewTextStyled(text *StyledString, width, height int) *Text {
	return nil
}

// Render renders the text onto the given screen.
func (c *Text) Render(s *Screen, e Event) error {
	return nil
}

// Text returns the text.
func (c *Text) Text() string {
	return c.text.String()
}

// SetText sets the text.
func (c *Text) SetText(text string) {
	c.text = NewStyledString(text)
}

// TextStyled returns the text with styles.
func (c *Text) TextStyled() *StyledString {
	return c.text
}

// SetTextStyled sets the text with styles.
func (c *Text) SetTextStyled(text *StyledString) {
	c.text = text
}

// Width returns the preferred width.
func (c *Text) Width() int {
	return 0
}

// SetWidth sets the preferred width of the box.
func (c *Text) SetWidth(width int) {
	c.width = width
}

// Height returns the preferred height.
func (c *Text) Height() int {
	return 0
}

// SetHeight sets the preferred height of the box.
func (c *Text) SetHeight(height int) {
	c.height = height
}
