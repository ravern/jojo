package mojo

// Stack represents a stack of components.
type Stack struct {
	// Direction is the direction of the stack.
	Direction StackDirection

	components []Component
}

// StackDirection is the direction of the stack.
type StackDirection int

// Possible directions of the stack.
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
