package mojo

// Component represents a generic component to be rendered.
type Component interface {
	// Render renders onto the given screen.
	//
	// The width and height of the given screen might not always be the same
	// as the preferred width and height of the components. This is due to
	// collisions and space constraints. Therefore, the width and height of
	// the given screen should be taken as the width and height to render
	// to.
	Render(*Screen) error

	// Click responds to a click and returns whether a response was made.
	//
	// If a component does not respond to clicks or chooses not to respond,
	// then false should be returned. Components that contain other
	// components should foward to call to them if appropriate.
	Click(x, y int) bool

	// Width returns the preferred width.
	//
	// This is useful for situations like content wrapping or absolute
	// widths. If there is no preferred width, 0 should be returned.
	Width() int

	// Height returns the preferred height.
	//
	// This is useful for situations like content wrapping or absolute
	// widths. If there is no preferred height, 0 should be returned.
	Height() int
}
