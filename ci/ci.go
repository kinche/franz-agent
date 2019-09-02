package ci

import (
	"fmt"
	"os"
)

type CI interface {
	Get() Info
}

// Info holds CI information
type Info struct {
	Sha1   string
	Branch string
	CI     string
}

// Detect detects which CI platform this agent is running
func Detect() CI {
	if os.Getenv("CIRCLECI") == "true" {
		fmt.Println("[franz-agent] detected CI environment: circle")
		return NewCircle()
	} else if os.Getenv("TRAVIS") == "true" {
		fmt.Println("[franz-agent] detected CI environment: travis")
		return NewTravis()
	}
	return NewCircle()
}
