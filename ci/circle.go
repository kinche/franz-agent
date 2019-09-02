package ci

import "os"

// Circle struct
type Circle struct{}

// Get returns necessary information from Circle
func (c Circle) Get() Info {
	return Info{
		Sha1:   os.Getenv("CIRCLE_SHA1"),
		Branch: os.Getenv("CIRCLE_BRANCH"),
		CI:     "circle",
	}
}

// NewCircle is a constructor for Circle
func NewCircle() Circle {
	return Circle{}
}
