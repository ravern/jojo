package text

import "github.com/ravernkoh/mojo"

// Label represents a label containing some text.
type Label struct {
	HoriAlignment Alignment
	VertAlignment Alignment

	width  int
	height int

	text StyledString
}

// NewLabel creates a new box with the given text, width and height.
func NewLabel(text string, width, height int) *Label {
	return nil
}

// NewLabelStyled creates a new box with the given text with styles, width and
// height.
func NewLabelStyled(text *StyledString, width, height int) *Label {
	return nil
}

// Render renders the text onto the given screen.
func (c *Label) Render(s *mojo.Screen, e mojo.Event) error {
	return nil
}

// Text returns the text.
func (c *Label) Text() string {
	return c.text.String()
}

// SetText sets the text.
func (c *Label) SetText(text string) {
	c.text = NewStyledString(text)
}

// TextStyled returns the text with styles.
func (c *Label) TextStyled() StyledString {
	return c.text
}

// SetTextStyled sets the text with styles.
func (c *Label) SetTextStyled(text StyledString) {
	c.text = text
}

// Width returns the preferred width.
func (c *Label) Width() int {
	return 0
}

// SetWidth sets the preferred width of the box.
func (c *Label) SetWidth(width int) {
	c.width = width
}

// Height returns the preferred height.
func (c *Label) Height() int {
	return 0
}

// SetHeight sets the preferred height of the box.
func (c *Label) SetHeight(height int) {
	c.height = height
}
