package mojo

// Cell represents a single cell to be rendered to.
type Cell struct {
	Rune       rune
	Foreground Color
	Background Color
	Attributes []Attribute
}

// Color represents the color of a cell.
type Color int

// Possible colors of a cell.
const (
	ColorDefault Color = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

// Attribute represents some attributes of text.
type Attribute int

// Possible attributes of text.
const (
	AttributeBold Attribute = 1 << (iota + 9)
	AttributeUnderline
	AttributeReverse
)
