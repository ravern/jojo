package mojo

import "strings"

// StyledString represents a string with styles.
type StyledString []StyledStringSection

// StyledStringSection represents a section of a string with styles.
type StyledStringSection struct {
	String     string
	Foreground Color
	Background Color
	Attributes []Attribute
}

// NewStyledString creates a new string with styles wrapping the given string.
func NewStyledString(s string) StyledString {
	return []StyledStringSection{{String: s}}
}

func (s StyledString) String() string {
	var b strings.Builder
	for _, section := range s {
		b.WriteString(section.String)
	}
	return b.String()
}
