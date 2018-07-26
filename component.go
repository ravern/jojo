package mojo

import "time"

// Component represents a generic component to be rendered.
type Component interface {
	// Render renders onto the given screen.
	//
	// The width and height of the given screen might not always be the same
	// as the preferred width and height of the components. This is due to
	// collisions and space constraints. Therefore, the width and height of
	// the given screen should be taken as the width and height to render
	// to.
	//
	// The event passed is the event that triggered the render. This can be
	// anything from a time-based event, to a mouse click or key press.
	Render(*Screen, Event) error

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

// Event represents an event that triggered a render.
type Event interface {
	// Time is the time that the event occured.
	//
	// The time returned is sometimes used for event replaying/syncing.
	// Therefore, events can be sent out of order. Components simply have
	// to keep track of this timing if they depend on the time.
	Time() time.Time
}
