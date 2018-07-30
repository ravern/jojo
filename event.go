package mojo

import "time"

// Event represents an event that triggered a render.
type Event interface {
	// Time is the time that the event occured.
	//
	// The time returned is sometimes used for event replaying/syncing.
	// Therefore, events can be sent out of order. Components simply have
	// to keep track of this timing if they depend on the time.
	Time() time.Time
}
