package ci

import (
	"os"

	"github.com/kinche/franz-agent/pkg/git"
)

// Circle struct
type Circle struct{}

// NewCircle is a constructor for Circle
func NewCircle() Circle {
	return Circle{}
}

// Get returns necessary information from Circle
func (c Circle) Get() Info {
	return Info{
		Sha1:   c.GetSha1(),
		Branch: c.GetBranch(),
		Author: c.GetAuthor(),
		CI:     "circle",
	}
}

// GetAuthor gets the author of the latest commit
func (c Circle) GetAuthor() string {
	return git.GetAuthor()
}

// GetBranch tries to fetch the branch name from the env var first.
// if it can't, it run a git command to get it
func (c Circle) GetBranch() string {
	s := os.Getenv("CIRCLE_BRANCH")
	if s == "" {
		s = git.GetBranchName()
	}
	return s
}

// GetSha1 tries to fetch the sha1 from the env var first
// if it can't, it run a git command to get it
func (c Circle) GetSha1() string {
	s := os.Getenv("CIRCLE_SHA1")
	if s == "" {
		s = git.GetSha1()
	}
	return s
}
