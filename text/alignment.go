package text

// Alignment represents the alignment of some text.
//
// Text needs to be aligned both in the horizontal and vertical direction. The
// alignment in both directions can be generalised as going from leading to
// trailing. The horizontal direction goes from left to right while the vertical
// direction goes from top to bottom.
type Alignment int

// Possible alignments of text.
const (
	AlignmentLeading Alignment = iota
	AlignmentCenter
	AlignmentTrailing
)
