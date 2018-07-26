package mojo

// Stack represents a stack of components.
type Stack struct {
	Direction StackDirection

	components []Component
}

// StackDirection represents the direction of a stack.
type StackDirection int

// Possible directions of a stack.
const (
	StackDirectionVertical StackDirection = iota
	StackDirectionHorizontal
)

// Render renders the stack of child components.
func (c *Stack) Render(s *Screen) error {
	return nil
}

// Width returns the preferred width.
func (c *Stack) Width() int {
	return 0
}

// Height returns the preferred height.
func (c *Stack) Height() int {
	return 0
}
