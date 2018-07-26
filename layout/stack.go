package layout

import "github.com/ravernkoh/mojo"

// Stack represents a stack of components.
type Stack struct {
	Direction StackDirection

	components []mojo.Component
}

// StackDirection represents the direction of a stack.
type StackDirection int

// Possible directions of a stack.
const (
	StackDirectionVertical StackDirection = iota
	StackDirectionHorizontal
)

// Render renders the stack of child components.
func (c *Stack) Render(s *mojo.Screen) error {
	return nil
}

// Click responds to a click.
func (c *Stack) Click(x, y int) bool {
	return false
}

// Width returns the preferred width.
func (c *Stack) Width() int {
	return 0
}

// Height returns the preferred height.
func (c *Stack) Height() int {
	return 0
}
